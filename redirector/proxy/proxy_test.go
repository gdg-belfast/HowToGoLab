package proxy

import (
	"testing"
)

// Make sure our redirects are being set correctly
func TestSetRedirects(t *testing.T) {

	SetRedirects(map[string]string{
		"/test1": "https://else.com",
		"/test2": "https://tohere.com",
	})

	// make sure our values are being set correctly
	if to, ok := Redirects["/test1"]; ok {
		if to != "https://else.com" {
			t.Error("Failed to find correct value for /test1")
		}
	} else {
		t.Error("Failed to find /test1")
	}
	if to, ok := Redirects["/test2"]; ok {
		if to != "https://tohere.com" {
			t.Error("Failed to find correct value for /test2")
		}
	} else {
		t.Error("Failed to find /test2")
	}
}
