package utils

import (
	"fmt"
	"bytes"
	"strings"
	"net/http"
	"crypto/tls"
)

type SplunkManagement struct {
	DomainName string
	Port int
	Username string
	Password string
}

func (s *SplunkManagement) CreateIndex(index string) error {
	endpoint := fmt.Sprintf("https://%s:%d/%s", strings.TrimSuffix(s.DomainName, "/"), s.Port, "servicesNS/admin/search/data/indexes")
	data := []byte("name=" + index)

    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client := &http.Client{Transport: tr}
	request, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	request.SetBasicAuth(s.Username, s.Password)
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	res, error := client.Do(request) ; if error != nil {
		return err
	}

	defer res.Body.Close()

	switch res.StatusCode {
	case http.StatusOK:
		return nil
	case http.StatusCreated:
		return nil
	case http.StatusConflict:
		return nil
	case http.StatusUnauthorized:
		return fmt.Errorf("Could not create index %s: invalid password/username (%s:%s)", index, s.Username, s.Password)
	default:
		return fmt.Errorf("Could not create index %s: got error %d", index, res.StatusCode)
	}
}