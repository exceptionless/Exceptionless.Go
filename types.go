package exceptionless

import (
	"github.com/satori/go.uuid"
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

// GetBaseEvent of function type
type GetBaseEvent func(string, string, string) Event

// AddSource of function type
type AddSource func(Event, string) Event

// AddTags of function type
type AddTags func(Event, []string) Event

// AddGeo of function type
type AddGeo func(Event, string) Event

// AddValue of function type
type AddValue func(Event, uint) Event

// AddReferenceID of function type
type AddReferenceID func(Event, uuid.UUID) Event

// AddCount of function type
type AddCount func(Event, uint) Event

// AddLogLevel of function type
type AddLogLevel func(Event, string) Event

// AddData of function type
type AddData func(Event, map[string]interface{}) Event

// SubmitError of function type
type SubmitError func(error) string

// SubmitLog of function type
type SubmitLog func(string, string) string

// SubmitEvent of function type
type SubmitEvent func(string) string