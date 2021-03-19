package main

import (
	"fmt"
	"math/rand"
	"time"
)

var config map[string]interface{} = nil

//Exceptionless type defines the client configuration structure
type Exceptionless struct {
	apiKey                         string
	serverURL                      string
	updateSettingsWhenIdleInterval int32
	includePrivateInformation      bool
}

var client = Exceptionless{}

func main() {
	handlePolling()
}

func handlePolling() {
	if client.apiKey != "" && client.updateSettingsWhenIdleInterval > 0 {
		fmt.Println("polling!")
		poll()
	}
}

//Configure is the function that creates an Exceptionless client
func Configure(config Exceptionless) Exceptionless {
	client = config
	handlePolling()
	return client
}

//GetClient returns the Exceptionless client
func GetClient() Exceptionless {
	return client
}

//GetConfig returns the project configuration
func GetConfig() map[string]interface{} {
	return config
}

//SubmitEvent sends log events to Exceptionless
func SubmitEvent(eventBody string) string {
	if client.apiKey == "" {
		fmt.Println("is zero value")
	}
	resp := Post("events", eventBody, client.apiKey)
	return resp
}

func poll() {
	r := rand.New(rand.NewSource(99))
	c := time.Tick(10 * time.Second)
	for _ = range c {
		resp := Get("projects/config", client.apiKey)
		config = resp
		jitter := time.Duration(r.Int31n(client.updateSettingsWhenIdleInterval)) * time.Millisecond
		time.Sleep(jitter)
	}
}
