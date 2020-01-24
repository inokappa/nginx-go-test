package main

import (
	_ "fmt"
	"io/ioutil"
	"net/http"
	"os"
	_ "strings"
	"testing"
)

const (
	CODE    = 400
	MESSAGE = "Bad Request!!!"
)

var (
	URL = os.Getenv("URL")
)

func TestApiV4(t *testing.T) {
	urls := []string{
		URL + "/foo/123",
		URL + "/bar/456",
		URL + "/baz/789/000",
	}

	runTests(urls, t)
}

func runTests(urls []string, t *testing.T) {
	for _, url := range urls {
		t.Run("URL: "+url, func(t *testing.T) {
			// Build Request
			req, _ := http.NewRequest("GET", url, nil)
			// Execute Request
			res, err := executeRequest(req)
			if err != nil {
				t.Fatalf("http.Get: %v", err)
			}
			defer res.Body.Close()

			// Display Test Result
			testResults(res, t)
		})
	}
}

func executeRequest(req *http.Request) (*http.Response, error) {
	client := new(http.Client)
	// Disable redirect following
	// Go の http package はデフォルトで自動的にリダイレクトを追っかける
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func testResults(res *http.Response, t *testing.T) {
	// Status Code Test
	if res.StatusCode != CODE {
		t.Errorf("got %d, want %d", res.StatusCode, CODE)
	}

	// Request Header Test (Nginx version number removed.)
	s := res.Header.Get("Server")
	if s != "nginx" {
		t.Errorf("got %s, want %s", s, "nginx")
	}

	// Response Body Test
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("Read Body Error: %v", err)
	}
	stringBody := string(b)

	if stringBody != MESSAGE {
		t.Errorf("got %s, want %s", stringBody, MESSAGE)
	}
}
