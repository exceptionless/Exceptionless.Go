package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

var config = ""

type Exceptionless struct {
	apiKey                         string
	serverURL                      string
	configServerURL                string
	heartbeatServerURL             string
	updateSettingsWhenIdleInterval uint
	includePrivateInformation      bool
}

var client = Exceptionless{}

func main() {
	if client.apiKey != "" {
		poll()
	}
	fmt.Println("No Exceptionless client defined")
}

func Configure(config Exceptionless) Exceptionless {
	client = config

	if client.apiKey != "" {
		poll()
	}
	return client
}

func poll() {
	r := rand.New(rand.NewSource(99))
	c := time.Tick(10 * time.Second)
	for _ = range c {
		url := "https://api.exceptionless.io/api/v2/projects/config"
		method := "GET"

		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)

		if err != nil {
			fmt.Println(err)
			return
		}
		req.Header.Add("accept", "application/json")
		req.Header.Add("Authorization", "Bearer XUlBBdgFxAlmCsAZHDFTIacXpzYuZDuqDzzFYMlR")

		res, err := client.Do(req)
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
