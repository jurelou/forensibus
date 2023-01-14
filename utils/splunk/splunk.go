package splunk

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/jurelou/forensibus/utils"
)

var (
	hecEndpoint               = "/services/collector"
	maxBufferSize             = 5000000 // 5MB
	timeout                   = 5 * time.Second
	retryWait                 = 1 * time.Second
	maxRetries                = 3
	statusServerBusy          = 9
	statusInternalServerError = 8
)

type Event struct {
	Host       string      `json:"host,omitempty"`
	Index      string      `json:"index,omitempty"`
	Source     string      `json:"source,omitempty"`
	SourceType string      `json:"sourcetype,omitempty"`
	Time       string      `json:"time,omitempty"`
	Event      interface{} `json:"event"`
}

func NewEvent(event interface{}) *Event {
	return &Event{Event: event}
}

type Response struct {
	Text  string          `json:"text"`
	Code  int             `json:"code"`
	AckID *int            `json:"ackId"`
	Acks  map[string]bool `json:"acks"`
}

type HEC struct {
	client *http.Client

	endpoint string
	token    string

	buffer bytes.Buffer

	defaultHost       string
	defaultIndex      string
	defaultSource     string
	defaultSourceType string
}

func NewHEC(address string, token string) HEC {
	transport := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true}, // allow insecure TLS
		MaxIdleConns:        10,                                    // global number of idle conns
		MaxIdleConnsPerHost: 10,                                    // subset of MaxIdleConns, per-host
		IdleConnTimeout:     2 * time.Second,
		// DisableKeepAlives: true, // this means create a new connection per request. not recommended
	}
	client := http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
	endpoint := strings.TrimSuffix(address, "/") + hecEndpoint
	return HEC{endpoint: endpoint, token: token, client: &client}
}

func (hec *HEC) Close() {
	hec.flushBuffer()
}

func (hec *HEC) SetDefaultIndex(index string) {
	hec.defaultIndex = index
}

func (hec *HEC) SetDefaultSource(source string) {
	hec.defaultSource = source
}

func (hec *HEC) SetDefaultSourceType(sourceType string) {
	hec.defaultSourceType = sourceType
}

func (hec *HEC) SetDefaultHost(host string) {
	hec.defaultHost = host
}

func (hec *HEC) makeRequest(data bytes.Buffer) error {
	retries := 0

	// gzip compress buffer
	var compressedBuffer bytes.Buffer
	gzipWriter := gzip.NewWriter(&compressedBuffer)
	_, err := gzipWriter.Write(data.Bytes())
	gzipWriter.Close()
	if err != nil {
		return err
	}

LOOP:
	if retries > maxRetries {
		return fmt.Errorf("Tried %d times to send splunk data, exiting.", retries)
	}

	req, err := http.NewRequest(http.MethodPost, hec.endpoint, &compressedBuffer)
	if err != nil {
		return err
	}
	// req = req.WithContext(ctx)
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Authorization", "Splunk "+hec.token)
	req.Header.Set("Content-Encoding", "gzip")

	res, err := hec.client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		var resBody Response
		if err := json.Unmarshal(body, &resBody); err != nil {
			return fmt.Errorf("Invalid response from splunk: %s", string(body))
		}

		if resBody.Code == statusServerBusy || resBody.Code == statusInternalServerError {
			utils.Log.Infof("Error while sending data to splunk, retrying (%d/%d)", retries, maxRetries)
			retries++
			time.Sleep(retryWait)
			goto LOOP
		}
	}
	utils.Log.Debugf("Sent %d bytes to splunk", hec.buffer.Len())

	return nil
}

func (hec *HEC) flushBuffer() {
	if hec.buffer.Len() == 0 {
		return
	}
	if err := hec.makeRequest(hec.buffer); err != nil {
		utils.Log.Warnf("Error while sending %d bytes to splunk: %s", hec.buffer.Len(), err.Error())
	}
	hec.buffer.Reset()
}

func (hec *HEC) WriteEvent(event *Event) {
	if event.Event == nil {
		utils.Log.Warnf("Tried to send empty event to splunk")
		return
	}
	if event.Index == "" {
		event.Index = hec.defaultIndex
	}
	if event.Source == "" {
		event.Source = hec.defaultSource
	}
	if event.SourceType == "" {
		event.SourceType = hec.defaultSourceType
	}
	if event.Host == "" {
		event.Host = hec.defaultHost
	}

	json, _ := json.Marshal(event)

	if len(json)+hec.buffer.Len() > maxBufferSize {
		// Size exceed max buffer size, flush buffer
		hec.flushBuffer()
	}
	hec.buffer.Write(json)
	hec.buffer.WriteString("\n")
}

// type SplunkManagement struct {
// 	DomainName string
// 	Port int
// 	Username string
// 	Password string
// }

// func (s *SplunkManagement) CreateIndex(index string) error {
// 	endpoint := fmt.Sprintf("https://%s:%d/%s", strings.TrimSuffix(s.DomainName, "/"), s.Port, "servicesNS/admin/search/data/indexes")
// 	data := []byte("name=" + index)

//     tr := &http.Transport{
//         TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
//     }
//     client := &http.Client{Transport: tr}
// 	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
// 	if err != nil {
// 		return err
// 	}
// 	request.SetBasicAuth(s.Username, s.Password)
// 	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
// 	res, error := client.Do(request) ; if error != nil {
// 		return err
// 	}

// 	defer res.Body.Close()

// 	switch res.StatusCode {
// 	case http.StatusOK:
// 		return nil
// 	case http.StatusCreated:
// 		return nil
// 	case http.StatusConflict:
// 		return nil
// 	case http.StatusUnauthorized:
// 		return fmt.Errorf("Could not create index %s: invalid password/username (%s:%s)", index, s.Username, s.Password)
// 	default:
// 		return fmt.Errorf("Could not create index %s: got error %d", index, res.StatusCode)
// 	}
// }
