package main

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

//Post posts to the Exceptionless Server
func Post(endpoint string, postBody string, authorization string) string {
	println(postBody)
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
