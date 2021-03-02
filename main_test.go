package main

import (
	"fmt"
	"testing"
)

// func TestClientConfiguration(t *testing.T) {
// 	err := godotenv.Load(".env")

// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}

// 	// getting env variables SITE_TITLE and DB_HOST
// 	testKey := os.Getenv("EXCEPTIONLESS_TEST_KEY")
// 	var settings Exceptionless
// 	settings.apiKey = testKey
// 	var client Exceptionless = Configure(settings)
// 	if client.apiKey == "" {
// 		t.Errorf("is zero value")
// 	}
// }

// func TestGetProjectConfig(t *testing.T) {
// 	config := GetProjectConfig()
// 	if config == "" {
// 		t.Errorf("No project config returned")
// 	}
// 	fmt.Println(config)
// }

func TestBuildSimpleLog(t *testing.T) {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// // getting env variables SITE_TITLE and DB_HOST
	// testKey := os.Getenv("EXCEPTIONLESS_TEST_KEY")
	// var settings Exceptionless
	// settings.apiKey = testKey
	// var client Exceptionless = Configure(settings)
	// if client.apiKey == "" {
	// 	t.Errorf("is zero value")
	// }
	// log := `{ "type": "log", "message": "Exceptionless is amazing!", "date":"2030-01-01T12:00:00.0000000-05:00", "@user":{ "identity":"123456789", "name": "Test User" } }`
	// // SubmitLog(log)
	// SubmitLog("line36 app.js", "Undefined is not a value", "error")
	var event Event
	event = GetBaseEvent("log", "boom son")
	event = AddSource(event, "line 66 app.js")
	fmt.Println(event)
}

func TestBuildLogWithTags(t *testing.T) {
	// err := godotenv.Load(".env")

	// if err != nil {
	// 	log.Fatalf("Error loading .env file")
	// }

	// // getting env variables SITE_TITLE and DB_HOST
	// testKey := os.Getenv("EXCEPTIONLESS_TEST_KEY")
	// var settings Exceptionless
	// settings.apiKey = testKey
	// var client Exceptionless = Configure(settings)
	// if client.apiKey == "" {
	// 	t.Errorf("is zero value")
	// }
	// log := `{ "type": "log", "message": "Exceptionless is amazing!", "date":"2030-01-01T12:00:00.0000000-05:00", "@user":{ "identity":"123456789", "name": "Test User" } }`
	// // SubmitLog(log)
	// SubmitLog("line36 app.js", "Undefined is not a value", "error")
	var event Event
	event = GetBaseEvent("log", "boom son")
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	fmt.Println(event)
}
