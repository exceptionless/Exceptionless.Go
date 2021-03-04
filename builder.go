package main

type Event struct {
	eventType   string
	source      string
	date        string
	tags        []string
	message     string
	geo         string
	value       uint
	data        map[string]interface{}
	referenceID string
	count       uint
}

//GetBaseEvent returns an empty Event struct that can be built into any type of event.
func GetBaseEvent(eventType string, message string, date string) Event {
	return Event{
		eventType: eventType,
		message:   message,
		date:      date,
	}
}

//AddSource adds a string value source to an event
func AddSource(event Event, source string) Event {
	event.source = source
	return event
}

//AddTags adds a string array of tags for the event
func AddTags(event Event, tags []string) Event {
	event.tags = tags
	return event
}

//AddGeo adds the lat and long location of the user the event impacted
func AddGeo(event Event, geo string) Event {
	event.geo = geo
	return event
}

//AddValue adds an arbitrary number value to the event
func AddValue(event Event, value uint) Event {
	event.value = value
	return event
}

//AddReferenceID adds an indentifier to later refer to this event
func AddReferenceID(event Event, referenceID string) Event {
	event.referenceID = referenceID
	return event
}

//AddCount adds a number to help track the number of times the event has occurred
func AddCount(event Event, count uint) Event {
	event.count = count
	return event
}

//AddData adds a string mapping to create a data object of additional values
func AddData(event Event, data map[string]interface{}) Event {
	event.data = data
	return event
}

// {
//   "type": "error",
//   "source": "Website", // Where the event came from in the app (something to stack on)
//   "reference_id": "123",
//   "message": "some event message",
//   "geo": `${latitude},${longitude}`,
//   "date":"2030-01-01T12:00:00.0000000-05:00",
//   "data": {
//     "@ref": {
//       "id": "parent event reference id",
//       "name": "parent event reference name"
//     },
//     "@user": {
//       "identity": "email or something",
//       "name": "John Doe",
//       "data": "Anything we want"
//     },
//     "@user_description": {
//       "email_address": "email",
//       "description": "super cool user",
//       "data": "Anything we want"
//     },
//     "@stack": { //  If provided, changes the default stacking mannerism and forces stacking based on info passed here.
//       "signature_data": {
//         "ManualStackingKey": "manual key we set"
//       },
//       "title": "stack title"
//     },
//   },
//   "value": "some number",  //Int representing anything
//   "tags": ["string", "string", "string"]
// }
