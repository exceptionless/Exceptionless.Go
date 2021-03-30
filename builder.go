package exceptionless

import (
	"encoding/json"
	"fmt"
	"time"
	"github.com/satori/go.uuid"
	"github.com/go-errors/errors"
)

//Event is the main model for events
type Event struct {
	EventType   string                 `json:"type"`
	Source      string                 `json:"source,omitempty"`
	Date        string                 `json:"date,omitempty"`
	Tags        []string               `json:"tags,omitempty"`
	Message     string                 `json:"message,omitempty"`
	Geo         string                 `json:"geo,omitempty"`
	Value       uint                   `json:"value,omitempty"`
	Data        map[string]interface{} `json:"data,omitempty"`
	ReferenceID uuid.UUID                 `json:"reference_id,omitempty"`
	Count       uint                   `json:"count,omitempty"`
}

func GetBaseURL() string {
	return "https://collector.exceptionless.com/api/v2/"
}

//GetBaseEvent returns an empty Event struct that can be built into any type of event.
func GetBaseEvent(eventType string, message string, date string) Event {
	return Event{
		EventType: eventType,
		Message:   message,
		Date:      date,
	}
}

//AddSource adds a string value source to an event
func AddSource(event Event, source string) Event {
	event.Source = source
	return event
}

//AddTags adds a string array of tags for the event
func AddTags(event Event, tags []string) Event {
	event.Tags = tags
	return event
}

//AddGeo adds the lat and long location of the user the event impacted
func AddGeo(event Event, geo string) Event {
	event.Geo = geo
	return event
}

//AddValue adds an arbitrary number value to the event
func AddValue(event Event, value uint) Event {
	event.Value = value
	return event
}

//AddReferenceID adds an indentifier to later refer to this event
func AddReferenceID(event Event, referenceID uuid.UUID) Event {
	event.ReferenceID = referenceID
	return event
}

//AddCount adds a number to help track the number of times the event has occurred
func AddCount(event Event, count uint) Event {
	event.Count = count
	return event
}

func AddLogLevel(event Event, level string) Event {
	var updatedEvent Event
	if event.Data != nil {
		event.Data["@level"] = level
		updatedEvent = event
	} else {
		data := map[string]interface{}{}
		data["@level"] = level
		updatedEvent = AddData(event, data)
	}		
	return updatedEvent
}

//AddData adds a string mapping to create a data object of additional values
func AddData(event Event, data map[string]interface{}) Event {
	event.Data = data
	return event
}

//SubmitError is a convenience wrapper to quickly build and submit an error
func SubmitError(err error) string {
	if ExceptionlessClient.UpdateSettingsWhenIdleInterval > 0 {
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
	resp := SubmitEvent(string(json))
	return resp
}

func SubmitLog(message string, level string) string {
	exceptionlessClient := GetClient()
	referenceID := uuid.Must(uuid.NewV4())
	if exceptionlessClient.UpdateSettingsWhenIdleInterval > 0 {
		config := GetConfig()		
		fmt.Println(config)
		//	We are stripping out info accoring to the config settings
		//	We would also prevent logs of levels below the log level set by the settings
	}
	var event Event
	date := time.Now().Format(time.RFC3339)
	event = GetBaseEvent("log", message, date)
	event = AddReferenceID(event, referenceID)
	data := map[string]interface{}{}
	data["@level"] = level
	event = AddData(event, data)

	json, err := json.Marshal(event)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	resp := SubmitEvent(string(json))
	return resp
}

func SubmitEvent(eventBody string) string {
	if ExceptionlessClient.ApiKey == "" {
		fmt.Println("is zero value")
	}
	resp := Post("events", eventBody, ExceptionlessClient.ApiKey)
	return resp
}