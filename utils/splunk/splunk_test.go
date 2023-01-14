package splunk_test

import (
	// "strings"
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"bytes"
	"compress/gzip"
	// "github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/splunk"
)

func TestCreateEvent(t *testing.T) {
    svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Invalid http verb: %s", r.Method)
		}
		if r.URL.String() != "/services/collector" {
			t.Errorf("Invalid url: %s", r.URL)
		}
		if r.Header.Get("Authorization") != "Splunk test-token" {
			t.Errorf("Invalid auth token: %s", r.Header.Get("Authorization"))
		}
		bodyCompressed, err := ioutil.ReadAll(r.Body) ; if err != nil {
			t.Errorf("Could not read body: %s", err.Error())
		}
		reader := bytes.NewReader([]byte(bodyCompressed))
		gzipReader, err:= gzip.NewReader(reader); if err != nil {
			t.Errorf("Could not decompress body: %s", err.Error())
		}
		body, err := ioutil.ReadAll(gzipReader); if err != nil {
			t.Errorf("Could not read decompressed body: %s", err.Error())
		}
		if string(body) != "{\"host\":\"host_x\",\"index\":\"main_x\",\"source\":\"source_x\",\"sourcetype\":\"stype_x\",\"event\":\"test_event\"}\n" {
			t.Errorf("Invalid body: %s", string(body))
		}
		fmt.Fprintf(w, "some return")
    }))
    defer svr.Close()

	hec := splunk.NewHEC(svr.URL, "test-token")
	defer hec.Close()
	hec.SetDefaultHost("host_x")
	hec.SetDefaultIndex("main_x")
	hec.SetDefaultSource("source_x")
	hec.SetDefaultSourceType("stype_x")

	e := splunk.NewEvent("test_event")
	hec.WriteEvent(e)
}

func TestCreateEventNoFlush(t *testing.T) {
    svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("HEC should not be called (buffer not flushed")
    }))
    defer svr.Close()

	hec := splunk.NewHEC(svr.URL, "test-token")
	e := splunk.NewEvent("test_event")
	hec.WriteEvent(e)
}

func TestCreateEmptyEvent(t *testing.T) {
    svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("HEC should not be called (empty buffer)")
	}))
    defer svr.Close()

	hec := splunk.NewHEC(svr.URL, "test-token")
	defer hec.Close()
	e := splunk.NewEvent(nil)
	hec.WriteEvent(e)
}

func TestCreateEventOverride(t *testing.T) {
    svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			t.Errorf("Invalid http verb: %s", r.Method)
		}
		if r.URL.String() != "/services/collector" {
			t.Errorf("Invalid url: %s", r.URL)
		}
		if r.Header.Get("Authorization") != "Splunk test-token" {
			t.Errorf("Invalid auth token: %s", r.Header.Get("Authorization"))
		}
		bodyCompressed, err := ioutil.ReadAll(r.Body) ; if err != nil {
			t.Errorf("Could not read body: %s", err.Error())
		}
		reader := bytes.NewReader([]byte(bodyCompressed))
		gzipReader, err:= gzip.NewReader(reader); if err != nil {
			t.Errorf("Could not decompress body: %s", err.Error())
		}
		body, err := ioutil.ReadAll(gzipReader); if err != nil {
			t.Errorf("Could not read decompressed body: %s", err.Error())
		}
		if string(body) != "{\"host\":\"new_host\",\"index\":\"new_index\",\"source\":\"new_source\",\"sourcetype\":\"new_stype\",\"time\":\"new_time\",\"event\":\"new_event\"}\n" {
			t.Errorf("Invalid body: %s", string(body))
		}
    }))
    defer svr.Close()

	hec := splunk.NewHEC(svr.URL, "test-token")
	defer hec.Close()
	hec.SetDefaultHost("host_x")
	hec.SetDefaultIndex("main_x")
	hec.SetDefaultSource("source_x")
	hec.SetDefaultSourceType("stype_x")

	e := splunk.NewEvent("new_event")
	e.Host = "new_host"
	e.Index = "new_index"
	e.Source = "new_source"
	e.SourceType = "new_stype"
	e.Time = "new_time"

	hec.WriteEvent(e)
}