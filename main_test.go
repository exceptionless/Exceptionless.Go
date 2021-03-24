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
	Configure(settings)
	if ExceptionlessClient.apiKey == "" {
		t.Errorf("is zero value")
	}
}

func TestBuildSimpleEvent(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = ExceptionlessClient.getBaseEvent("log", "boom son", date)
	event = ExceptionlessClient.addSource(event, "line 66 app.js")
	if event.Source == "" {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithTags(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = ExceptionlessClient.getBaseEvent("log", "boom son", date)
	event = ExceptionlessClient.addSource(event, "line 66 app.js")
	event = ExceptionlessClient.addTags(event, []string{"one", "two", "three"})
	if event.Tags == nil {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithGeo(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = ExceptionlessClient.getBaseEvent("log", "boom son", date)
	event = ExceptionlessClient.addSource(event, "line 66 app.js")
	event = ExceptionlessClient.addTags(event, []string{"one", "two", "three"})
	event = ExceptionlessClient.addGeo(event, "44.14561, -172.32262")
	if event.Geo == "" {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithLogLevel(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = ExceptionlessClient.getBaseEvent("log", "boom son", date)
	event = ExceptionlessClient.addSource(event, "line 66 app.js")
	event = ExceptionlessClient.addTags(event, []string{"one", "two", "three"})
	event = ExceptionlessClient.addGeo(event, "44.14561, -172.32262")
	event = ExceptionlessClient.addLogLevel(event, "info")
	if event.Geo == "" {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithValue(t *testing.T) {
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = ExceptionlessClient.getBaseEvent("log", "boom son", date)
	event = ExceptionlessClient.addSource(event, "line 66 app.js")
	event = ExceptionlessClient.addTags(event, []string{"one", "two", "three"})
	event = ExceptionlessClient.addGeo(event, "44.14561, -172.32262")
	event = ExceptionlessClient.addValue(event, 21)
	if event.Value != 21 {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithReferenceID(t *testing.T) {
	var event Event
	referenceID := uuid.Must(uuid.NewV4())
	date := time.Now().Format(time.RFC3339)
	event = ExceptionlessClient.getBaseEvent("log", "boom son", date)
	event = ExceptionlessClient.addSource(event, "line 66 app.js")
	event = ExceptionlessClient.addTags(event, []string{"one", "two", "three"})
	event = ExceptionlessClient.addGeo(event, "44.14561, -172.32262")
	event = ExceptionlessClient.addValue(event, 21)
	event = ExceptionlessClient.addReferenceID(event, referenceID)
	if event.Source == "" {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithCount(t *testing.T) {
	var event Event
	referenceID := uuid.Must(uuid.NewV4())
	date := time.Now().Format(time.RFC3339)
	event = ExceptionlessClient.getBaseEvent("log", "boom son", date)
	event = ExceptionlessClient.addSource(event, "line 66 app.js")
	event = ExceptionlessClient.addTags(event, []string{"one", "two", "three"})
	event = ExceptionlessClient.addGeo(event, "44.14561, -172.32262")
	event = ExceptionlessClient.addValue(event, 21)
	event = ExceptionlessClient.addReferenceID(event, referenceID)
	event = ExceptionlessClient.addCount(event, 99)
	if event.Count != 99 {
		t.Errorf("Test failed")
	}
}

func TestBuildEventWithData(t *testing.T) {
	var event Event
	referenceID := uuid.Must(uuid.NewV4())
	date := time.Now().Format(time.RFC3339)
	event = ExceptionlessClient.getBaseEvent("log", "boom son", date)
	event = ExceptionlessClient.addSource(event, "line 66 app.js")
	event = ExceptionlessClient.addTags(event, []string{"one", "two", "three"})
	event = ExceptionlessClient.addGeo(event, "44.14561, -172.32262")
	event = ExceptionlessClient.addValue(event, 21)
	event = ExceptionlessClient.addReferenceID(event, referenceID)
	event = ExceptionlessClient.addCount(event, 99)
	e := map[string]interface{}{}
	e["message"] = "Whoops, an error"
	e["type"] = "System.Exception"
	e["stack_trace"] = " at Client.Tests.ExceptionlessClientTests.CanSubmitSimpleException() in ExceptionlessClientTests.cs:line 77"
	data := map[string]interface{}{}
	data["@error"] = e
	event = ExceptionlessClient.addData(event, data)
	if event.Data == nil {
		t.Errorf("Test failed")
	}
}

func TestErrorEvent(t *testing.T) {
	var event Event
	referenceID := uuid.Must(uuid.NewV4())
	date := time.Now().Format(time.RFC3339)
	event = ExceptionlessClient.getBaseEvent("error", "testing", date)
	event = ExceptionlessClient.addSource(event, "line 206 app.js")
	event = ExceptionlessClient.addTags(event, []string{"one", "two", "three"})
	event = ExceptionlessClient.addGeo(event, "44.14561, -172.32262")
	event = ExceptionlessClient.addValue(event, 21)
	event = ExceptionlessClient.addReferenceID(event, referenceID)
	event = ExceptionlessClient.addCount(event, 99)
	e := map[string]interface{}{}
	e["message"] = "Whoops, another"
	e["type"] = "System.Exception"
	e["stack_trace"] = " at Client.Tests.ExceptionlessClientTests.CanSubmitSimpleException() in ExceptionlessClientTests.cs:line 77"
	data := map[string]interface{}{}
	data["@error"] = e
	event = ExceptionlessClient.addData(event, data)
	json, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		t.Errorf("Test failed")
	}
	if string(json) == "" {
		t.Errorf("Test failed")
	}
	resp := ExceptionlessClient.submitEvent(string(json))
	if resp == "" {
		t.Errorf("Test failed")
	}
}

func TestSubmitError(t *testing.T) {
	e := errors.New(fmt.Sprintf("This is another error"))
	resp := ExceptionlessClient.submitError(e)
	if resp == "" {
		t.Errorf("Test failed")
	}
}

func TestSubmitInfoLog(t *testing.T) {
	message := "Info log!"
	level := "info"
	resp := ExceptionlessClient.submitLog(message, level)
	if resp == "" {
		t.Errorf("Test failed")
	}
}

func TestSubmitWarnLog(t *testing.T) {
	message := "Warn log!"
	level := "warn"
	resp := ExceptionlessClient.submitLog(message, level)
	if resp == "" {
		t.Errorf("Test failed")
	}
}
