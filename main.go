package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var config = ""

//Exceptionless type defines the client configuration structure
type Exceptionless struct {
	apiKey                         string
	serverURL                      string
	configServerURL                string
	heartbeatServerURL             string
	updateSettingsWhenIdleInterval uint
	includePrivateInformation      bool
	pollForConfig                  bool
}

var client = Exceptionless{}

func main() {
	if client.apiKey != "" && client.pollForConfig {
		poll()
	}
}

//Configure is the function that creates an Exceptionless client
func Configure(config Exceptionless) Exceptionless {
	client = config

	if client.apiKey != "" && client.pollForConfig {
		poll()
	}
	return client
}

//GetProjectConfig returns the project configuration settings
func GetProjectConfig() string {
	authorization := fmt.Sprintf("Bearer %s", client.apiKey)
	url := "https://api.exceptionless.io/api/v2/projects/config"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", authorization)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	config = string(body)
	return config
}

//SubmitEvent sends log events to Exceptionless
func SubmitEvent(eventBody string) string {
	if client.apiKey == "" {
		fmt.Println("is zero value")
	}
	fmt.Println(eventBody)
	resp := Post("events", eventBody, client.apiKey)
	return resp
}

func poll() {
	r := rand.New(rand.NewSource(99))
	c := time.Tick(10 * time.Second)
	for _ = range c {
		url := "https://api.exceptionless.io/api/v2/projects/config"
		method := "GET"

		httpClient := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}

		authorization := fmt.Sprintf("Bearer %s", client.apiKey)
		req.Header.Add("accept", "application/json")
		req.Header.Add("Authorization", authorization)

		res, err := httpClient.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		config = string(body)
		fmt.Printf("Grab at %s\n", time.Now())
		// add a bit of jitter
		jitter := time.Duration(r.Int31n(5000)) * time.Millisecond
		time.Sleep(jitter)
		fmt.Println(config)
	}
}
