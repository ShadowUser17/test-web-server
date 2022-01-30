package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWebServer(t *testing.T) {
	var mux = http.NewServeMux()
	mux.HandleFunc("/", ApiHandler)

	var server = httptest.NewUnstartedServer(mux)
	server.Start()

	var client = server.Client()

	if resp, err := client.Get(server.URL); err != nil {
		t.Errorf("TestWebServerGet: %v\n", err)

	} else {
		t.Logf("TestWebServerGet: %s\n", resp.Status)
		resp.Body.Close()
	}

	if resp, err := client.Head(server.URL); err != nil {
		t.Errorf("TestWebServerHead: %v\n", err)

	} else {
		t.Logf("TestWebServerHead: %s\n", resp.Status)
		resp.Body.Close()
	}

	var data = strings.NewReader("Testing...")

	if resp, err := client.Post(server.URL, "text/plain", data); err != nil {
		t.Errorf("TestWebServerPost: %v\n", err)

	} else {
		t.Logf("TestWebServerPost: %s\n", resp.Status)
		resp.Body.Close()
	}

	if req, err := http.NewRequest(http.MethodPut, server.URL, data); err != nil {
		t.Errorf("TestWebServer_NewRequest: %v\n", err)

	} else {
		if resp, err := client.Do(req); err != nil {
			t.Errorf("TestWebServerPut: %v\n", err)

		} else {
			t.Logf("TestWebServerPut: %s\n", resp.Status)
			resp.Body.Close()
		}
	}

	if req, err := http.NewRequest(http.MethodDelete, server.URL, nil); err != nil {
		t.Errorf("TestWebServer_NewRequest: %v\n", err)

	} else {
		if resp, err := client.Do(req); err != nil {
			t.Errorf("TestWebServerDelete: %v\n", err)

		} else {
			t.Logf("TestWebServerDelete: %s\n", resp.Status)
			resp.Body.Close()
		}
	}

	server.Close()
}
