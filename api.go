package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"fmt"
	"io/ioutil"
	"github.com/joho/godotenv"
	"encoding/json"
)

//Post posts to the Exceptionless Server
func Post(endpoint string, postBody string, authorization string) string {	
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables BASE_URL
	baseURL := os.Getenv("BASE_API_URL")
	url := baseURL + endpoint
	var jsonStr = []byte(postBody)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+authorization)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	// body, _ := ioutil.ReadAll(resp.Body)
	return string(resp.Status)
}

//GET makes api GET requests
func Get(endpoint string, authorization string) map[string]interface{} {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables BASE_URL
	baseURL := os.Getenv("BASE_API_URL")
	url := baseURL + endpoint

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		// return "Error"
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("Authorization", "Bearer "+authorization)

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		// return "Error"
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		// return "Error"
	}

	// resp := string(body)
	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	return result
}
