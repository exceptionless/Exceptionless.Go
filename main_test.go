package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing"
	"time"
	"github.com/satori/go.uuid"
	"github.com/go-errors/errors"
	"github.com/joho/godotenv"
)

func TestConfigureClient(t *testing.T) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables SITE_TITLE and DB_HOST
	testKey := os.Getenv("EXCEPTIONLESS_TEST_KEY")
	var settings Exceptionless
	settings.apiKey = testKey
	// settings.updateSettingsWhenIdleInterval = 3000 //This will enable polling for config
	var client Exceptionless = Configure(settings)
	if client.apiKey == "" {
		t.Errorf("is zero value")
	}
}

func TestBuildSimpleEvent(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	if event.Source == "" {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithTags(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	if event.Tags == nil {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithGeo(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	event = AddGeo(event, "44.14561, -172.32262")
	if event.Geo == "" {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithValue(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	event = AddGeo(event, "44.14561, -172.32262")
	event = AddValue(event, 21)
	if event.Value != 21 {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithReferenceID(t *testing.T) {
	var event Event
	referenceID := uuid.Must(uuid.NewV4())
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	event = AddGeo(event, "44.14561, -172.32262")
	event = AddValue(event, 21)
	event = AddReferenceID(event, referenceID)
	if event.Source == "" {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithCount(t *testing.T) {
	var event Event
	referenceID := uuid.Must(uuid.NewV4())
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	event = AddGeo(event, "44.14561, -172.32262")
	event = AddValue(event, 21)
	event = AddReferenceID(event, referenceID)
	event = AddCount(event, 99)
	if event.Count != 99 {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithData(t *testing.T) {
	var event Event
	referenceID := uuid.Must(uuid.NewV4())
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", "boom son", date)
	event = AddSource(event, "line 66 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	event = AddGeo(event, "44.14561, -172.32262")
	event = AddValue(event, 21)
	event = AddReferenceID(event, referenceID)
	event = AddCount(event, 99)
	e := map[string]interface{}{}
	e["message"] = "Whoops, an error"
	e["type"] = "System.Exception"
	e["stack_trace"] = " at Client.Tests.ExceptionlessClientTests.CanSubmitSimpleException() in ExceptionlessClientTests.cs:line 77"
	data := map[string]interface{}{}
	data["@error"] = e
	event = AddData(event, data)
	if event.Data == nil {
		t.Errorf("Test failed")
	}
}

func TestErrorEvent(t *testing.T) {
	var event Event
	referenceID := uuid.Must(uuid.NewV4())
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("error", "testing", date)
	event = AddSource(event, "line 206 app.js")
	event = AddTags(event, []string{"one", "two", "three"})
	event = AddGeo(event, "44.14561, -172.32262")
	event = AddValue(event, 21)
	event = AddReferenceID(event, referenceID)
	event = AddCount(event, 99)
	e := map[string]interface{}{}
	e["message"] = "Whoops, another"
	e["type"] = "System.Exception"
	e["stack_trace"] = " at Client.Tests.ExceptionlessClientTests.CanSubmitSimpleException() in ExceptionlessClientTests.cs:line 77"
	data := map[string]interface{}{}
	data["@error"] = e
	event = AddData(event, data)
	json, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		return
	}
	if string(json) == "" {
		t.Errorf("Test failed")
	}
	resp := SubmitEvent(string(json))
	fmt.Println(resp)
}

func TestSubmitException(t *testing.T) {
	e := errors.New(fmt.Sprintf("This is another error"))
	resp := SubmitException(e)
	if resp == "" {
		t.Errorf("Test failed")
	}
}
