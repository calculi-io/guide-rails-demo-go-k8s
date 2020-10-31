package main

import (
	// "fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"io/ioutil"

	"github.com/stretchr/testify/assert"
)

func TestHTTPHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(handler))
	defer ts.Close()

	http.HandleFunc("/", handler)
	r, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}
	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, string(response), "Congratulations! Your Go application has been successfully deployed on Kubernetes.")
}