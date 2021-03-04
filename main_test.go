package main

import (
	"fmt"
	"testing"
	"time"
)

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

func TestBuildSimpleLog(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	fmt.Println(event)
}

func TestBuildLogWithTags(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	fmt.Println(event)
}

func TestBuildLogWithGeo(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	event = AddGeo(event, "44.14561, -172.32262")
	fmt.Println(event)
}

func TestBuildLogWithValue(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	event = AddGeo(event, "44.14561, -172.32262")
	event = AddValue(event, 21)
	fmt.Println(event)
}

func TestBuildLogWithReferenceID(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	event = AddGeo(event, "44.14561, -172.32262")
	event = AddValue(event, 21)
	event = AddReferenceID(event, "11-22-33-44")
	fmt.Println(event)
}

func TestBuildLogWithCount(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	event = AddGeo(event, "44.14561, -172.32262")
	event = AddValue(event, 21)
	event = AddReferenceID(event, "11-22-33-44")
	event = AddCount(event, 99)
	fmt.Println(event)
}

func TestBuildLogWithData(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	event = AddGeo(event, "44.14561, -172.32262")
	event = AddValue(event, 21)
	event = AddReferenceID(event, "11-22-33-44")
	event = AddCount(event, 99)
	e := map[string]interface{}{}
	e["message"] = "Whoops, an error"
	e["type"] = "System.Exception"
	e["stack_trace"] = " at Client.Tests.ExceptionlessClientTests.CanSubmitSimpleException() in ExceptionlessClientTests.cs:line 77"
	data := map[string]interface{}{}
	data["@error"] = e
	event = AddData(event, data)
	fmt.Println(event)
}
