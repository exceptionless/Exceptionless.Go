package exceptionless

import (
	"fmt"
	"math/rand"
	"github.com/satori/go.uuid"
	"time"
	"encoding/json"
	"github.com/go-errors/errors"
)

var config map[string]interface{} = nil

//Exceptionless type defines the client configuration structure
type Exceptionless struct {
	apiKey                         string
	serverURL                      string
	updateSettingsWhenIdleInterval int32
	includePrivateInformation      bool
	getBaseEvent GetBaseEvent
	addSource AddSource
	addTags AddTags
	addGeo AddGeo
	addValue AddValue
	addReferenceID AddReferenceID
	addCount AddCount
	addLogLevel AddLogLevel
	addData AddData
	submitError SubmitError
	submitLog SubmitLog
	submitEvent SubmitEvent
}

//ExceptionlessClient returns the configured client
var ExceptionlessClient = Exceptionless{}

func handlePolling() {
	if ExceptionlessClient.apiKey != "" && ExceptionlessClient.updateSettingsWhenIdleInterval > 0 {
		fmt.Println("polling!")
		poll()
	}
}

//Configure is the function that creates an Exceptionless ExceptionlessClient
func Configure(config Exceptionless) Exceptionless {
	ExceptionlessClient = config
	ExceptionlessClient.getBaseEvent = func(eventType string, message string, date string) Event {
		return Event{
			EventType: eventType,
			Message:   message,
			Date:      date,
		}
	}
	ExceptionlessClient.addSource = func(event Event, source string) Event {
		event.Source = source
		return event
	}
	ExceptionlessClient.addTags = func(event Event, tags []string) Event {
		event.Tags = tags
		return event
	}
	ExceptionlessClient.addGeo = func(event Event, geo string) Event {
		event.Geo = geo
		return event
	}
	ExceptionlessClient.addValue = func(event Event, value uint) Event {
		event.Value = value
		return event
	}
	ExceptionlessClient.addReferenceID = func(event Event, referenceID uuid.UUID) Event {
		event.ReferenceID = referenceID
		return event
	}
	ExceptionlessClient.addCount = func(event Event, count uint) Event {
		event.Count = count
		return event
	}
	ExceptionlessClient.addLogLevel = func(event Event, level string) Event {
		var updatedEvent Event
		if event.Data != nil {
			event.Data["@level"] = level
			updatedEvent = event
		} else {
			data := map[string]interface{}{}
			data["@level"] = level
			updatedEvent = ExceptionlessClient.addData(event, data)
		}		
		return updatedEvent
	}
	ExceptionlessClient.addData = func(event Event, data map[string]interface{}) Event {
		event.Data = data
		return event
	}
	ExceptionlessClient.submitError = func(err error) string {
		exceptionlessClient := GetClient()
		if exceptionlessClient.updateSettingsWhenIdleInterval > 0 {
			config := GetConfig()		
			fmt.Println(config)
			//	We are stripping out info accoring to the config settings
		}
		referenceID := uuid.Must(uuid.NewV4())
		errorMap := map[string]interface{}{}
		errorMap["type"] = "error"
		errorMap["message"] = err.Error()
		errorMap["date"] = time.Now().Format(time.RFC3339)
		errorMap["stack_trace"] = err.(*errors.Error).ErrorStack()
		data := map[string]interface{}{}
		data["@simple_error"] = errorMap
		var mainEvent = Event{
			EventType: "error",
			Message: err.Error(),
			Data:      data,
			ReferenceID: referenceID,
		}
		json, err := json.Marshal(mainEvent)
		if err != nil {
			fmt.Println(err)
			return err.Error()
		}
		resp := ExceptionlessClient.submitEvent(string(json))
		return resp
	}
	ExceptionlessClient.submitLog = func(message string, level string) string {
		exceptionlessClient := GetClient()
		referenceID := uuid.Must(uuid.NewV4())
		if exceptionlessClient.updateSettingsWhenIdleInterval > 0 {
			config := GetConfig()		
			fmt.Println(config)
			//	We are stripping out info accoring to the config settings
			//	We would also prevent logs of levels below the log level set by the settings
		}
		var event Event
		date := time.Now().Format(time.RFC3339)
		event = ExceptionlessClient.getBaseEvent("log", message, date)
		event = ExceptionlessClient.addReferenceID(event, referenceID)
		data := map[string]interface{}{}
		data["@level"] = level
		event = ExceptionlessClient.addData(event, data)
	
		json, err := json.Marshal(event)
		if err != nil {
			fmt.Println(err)
			return err.Error()
		}
		resp := ExceptionlessClient.submitEvent(string(json))
		return resp
	}
	ExceptionlessClient.submitEvent = func(eventBody string) string {
		if ExceptionlessClient.apiKey == "" {
			fmt.Println("is zero value")
		}
		resp := Post("events", eventBody, ExceptionlessClient.apiKey)
		return resp
	}
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
		resp := Get("projects/config", ExceptionlessClient.apiKey)
		config = resp
		jitter := time.Duration(r.Int31n(ExceptionlessClient.updateSettingsWhenIdleInterval)) * time.Millisecond
		time.Sleep(jitter)
	}
}
