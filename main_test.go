package main

import "testing"

func TestClientConfiguration(t *testing.T) {
	var settings Exceptionless
	var client Exceptionless = Configure(settings)

	if client.apiKey == "" {
		t.Errorf("is zero value")
	}
}
