package exceptionless

import (
	"fmt"
	"math/rand"
	"time"
)

var config map[string]interface{} = nil

//Exceptionless type defines the client configuration structure
type Exceptionless struct {
	ApiKey                         string
	ServerURL                      string
	UpdateSettingsWhenIdleInterval int32
	IncludePrivateInformation      bool
}

//ExceptionlessClient returns the configured client
var ExceptionlessClient = Exceptionless{}

func handlePolling() {
	if ExceptionlessClient.ApiKey != "" && ExceptionlessClient.UpdateSettingsWhenIdleInterval > 0 {
		fmt.Println("polling!")
		poll()
	}
}

//Configure is the function that creates an Exceptionless ExceptionlessClient
func Configure(config Exceptionless) Exceptionless {
	ExceptionlessClient = config
	handlePolling()
	return ExceptionlessClient
}

//GetClient returns the Exceptionless client
func GetClient() Exceptionless {
	return ExceptionlessClient
}

//GetConfig returns the project configuration
func GetConfig() map[string]interface{} {
	return config
}

func poll() {
	r := rand.New(rand.NewSource(99))
	c := time.Tick(10 * time.Second)
	for _ = range c {
		resp := Get("projects/config", ExceptionlessClient.ApiKey)
		config = resp
		jitter := time.Duration(r.Int31n(ExceptionlessClient.UpdateSettingsWhenIdleInterval)) * time.Millisecond
		time.Sleep(jitter)
	}
}
