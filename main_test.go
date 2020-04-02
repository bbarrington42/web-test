package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	ts := httptest.NewUnstartedServer(new(PingHandler))
	ts.EnableHTTP2 = true
	ts.StartTLS()
	defer ts.Close()
	
	res, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Error(err)
	}
	
	pong, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		t.Error(err)
	}
	
	actual := string(pong)
	expected := "pong"
	
	if res.StatusCode != http.StatusOK {
		t.Errorf("StatusCode is %d, not %d", res.StatusCode, http.StatusOK)
	}
	
	if actual != expected {
		t.Errorf("%s does not equal %s", actual, expected)
	}
}
