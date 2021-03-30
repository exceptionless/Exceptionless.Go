# THIS IS A WORK IN PROGRESS - API SURFACE MAY CHANGE

# Exceptionless Go Client

This is a simple client wrapper to make working with Exceptionless in the Go environment easier. The client is not a complete wrapper around the Exceptionless API, so you are encouraged to review the [full Swagger documentation here](https://api.exceptionless.io/docs/index.html). 

## Getting Started 

To install the client, run `go get https://github.com/Exceptionless/Exceptionless.Go`.
 Import it into your project like this: 

```go
import (
	"github.com/Exceptionless/Exceptionless.Go"
)
```

Once you've imported it in your project, you'll need to configure the client. The current configuration options are as follows: 

* apiKey - string - required
* serverURL - string - optional

`apiKey` is self-explanatory. Get your key at https://exceptionless.com.

If you are self-hosting Exceptionless, provide the server URL for your self-hosted installation for the `serverURL` property. 

## Configuring Client

To set up your client, in the `main.go` file of your project, import Exceptionless as described above. 

There are two main ways to configure the Exceptionless client. You can set properties in the configuration struct one-by-one like this: 

```go
exceptionless.ExceptionlessClient.apiKey = "YOUR API KEY"
```

Or you can build up your own copy of the configuration struct and pass in your config settings all at once like this: 

```go
var settings = ExceptionlessClient
settings{
  apiKey: "YOUR API KEY"
  serverURL: "SELF HOSTED SERVER URL"
}

exceptionless.Configure(settings)
```

This will save your client information in-memory and will make it available throughout your app through a globally exposed variable called `ExceptionlessClient`. You can access that variable like this: `exceptionless.ExceptionlessClient`.

## Sending Events  

This client has two convenience functions: `SubmitError` and `SubmitLog`. These functions take minimal arguments and are far less flexible than building events yourself, but they are the easiest to use. We'll cover them first and then we will talk about how to build a custom event.

### SubmitError

The `SubmitError` function does exactly what it says. It will build a simple error event and submit it to Exceptionless. `SubmitError` takes a Go error type and returns a string value response. An example of how to call this function is in the `main_test.go` file as well as below: 

```go
func MyGoFunction() {
	e := errors.New(fmt.Sprintf("This is an error"))
	resp := exceptionless.SubmitError(e)
	fmt.Println(resp)
}
```

### SubmitLog

This function is a simple wrapper that allows you to submit log events with specified log levels. This function takes a `message` string and a `level` string. There is an example of this in `main_test.go` or below: 

```go
func MyGoFunction() {
	message := "Info log!"
	level := "info"
	resp := exceptionless.SubmitLog(message, level)
	fmt.Println(resp)
}
```

## Building Custom Events

This Go client allows you to build fully customized events. There are a number of helper functions to build these events. We'll walk through them all here. 

The first thing you will always need to do when building a custom event is get the base event that you will tack on additional info to. To do this, you would simple call the following function: 

```go
var event Event
date := time.Now().Format(time.RFC3339)
event = exceptionless.GetBaseEvent("log", "test log", date)
```

The `GetBaseEvent` function requires three parameters: eventType (string), message (string), and date (date).

### AddSource

This function will add a source location for the event. Think of this as an area of your application where the event happened. Here's an example of how to make use of this function: 

```go
var event Event
date := time.Now().Format(time.RFC3339)
event = exceptionless.GetBaseEvent("log", "test log", date)
event = exceptionless.AddSource(event, "line 66 main.go")
```

`AddSource` takes in your exist event as well as a string variable for the source of the event.

### AddTags

This function will add a string array of tags to your event. These tags can be used for filtering within the Exceptionless app. Here is an example: 

```go
var event Event
date := time.Now().Format(time.RFC3339)
event = exceptionless.GetBaseEvent("log", "test log", date)
event = exceptionless.AddTags(event, []string{"one", "two", "three"})
```

As you can see, the `AddTags` function takes the existing event and a string array as arguments. 

### AddGeo

This function will add the geographical location of your user with a latitude and logitude. Here is an example: 

```go
var event Event
date := time.Now().Format(time.RFC3339)
event = exceptionless.GetBaseEvent("log", "test log", date)
event = exceptionless.AddGeo(event, "44.14561, -172.32262")
```

The `AddGeo` function takes in your existing event and a string value with comma delimited lat and long values as arguments. 

### AddValue 

This function will add an arbitrary integer value to your events. It can represent anything you'd like. Here's an example: 

```go
var event Event
date := time.Now().Format(time.RFC3339)
event = exceptionless.GetBaseEvent("log", "test log", date)
event = exceptionless.AddValue(event, 21)
```

As you can see, the `AddValue` function takes in your existing event and an integer value as arguments. 

### AddReferenceID 

This function will add a UUID reference that you can use later to identify specific events. This identifier must be unique across your events.

**Note: When using the `SubmitEvent` and `SubmitLog` helpers, a referenceID will automatically be generated and applied to your events.**  

Here is an example of how to use this function: 

```go
var event Event
referenceID := uuid.Must(uuid.NewV4())
date := time.Now().Format(time.RFC3339)
event = exceptionless.GetBaseEvent("log", "test log", date)
event = exceptionless.AddReferenceID(event, referenceID)
```

As you can see, the `addReferenceId` function takes in your existing event and a referenceID (in the form of the UUID type) as arguments. 

### AddCount

This function adds a count integer to help you easily count just about anything you want. The count will be added to your event. Here is an example of how to use this function: 

```go
var event Event
date := time.Now().Format(time.RFC3339)
event = exceptionless.GetBaseEvent("log", "test log", date)
event = exceptionless.AddCount(event, 99)
```
As you can see, the `AddCount` function takes in your existing event and an integer as arguments. 

### AddLogLevel

This function is designed for log type events. It will add a log level to your event. This is useful for filtering. Here is an example of how to use this function: 

```go
var event Event
date := time.Now().Format(time.RFC3339)
event = exceptionless.GetBaseEvent("log", "test log", date)
event = exceptionless.AddLogLevel(event, "info")
```

As you can see, the `AddLogLevel` function takes in your existing event and a string value representing the log level as arguments. 

### AddData

Exceptionless can take any dictionary mapping of values that you want to provide. This helper function will apply your custom mapping for you. Here is an example of how to use this function: 

```go
var event Event
event = exceptionless.GetBaseEvent("log", "test log", date)
data := map[string]interface{}{}
data["custom_key"] = "custom string value"
exceptionless.AddData(event, data)
```

The data property takes a string map of arbitrary values. That said, there are some well-known keys that Exceptionless will treat as first-class citizens. [Read more about those here](https://exceptionless.com/docs/clients/custom-clients/#understanding-the-data-models).

### SubmitEvent

This is the function you call when your event has been built and is ready to be submitted to Exceptionless. Here is an example of how to use it: 

```go
json, err := json.Marshal(event)
if err != nil {
	fmt.Println(err)
}

resp := exceptionless.SubmitEvent(string(json))
```

## Full Example

Here is a more complete example of building and submitting a custom event: 

```go
var event Event
referenceID := uuid.Must(uuid.NewV4())
date := time.Now().Format(time.RFC3339)
event = exceptionless.GetBaseEvent("error", "testing", date)
event = exceptionless.AddSource(event, "line 206 app.js")
event = exceptionless.AddTags(event, []string{"one", "two", "three"})
event = exceptionless.AddGeo(event, "44.14561, -172.32262")
event = exceptionless.AddValue(event, 21)
event = exceptionless.AddReferenceID(event, referenceID)
event = exceptionless.AddCount(event, 99)
data := map[string]interface{}{}
data["custom_key"] = "custom string value"
event = exceptionless.AddData(event, data)
json, err := json.Marshal(event)
if err != nil {
	fmt.Println(err)
}
resp := exceptionless.SubmitEvent(string(json))
if resp == "" {
	fmt.Println("Test failed")
}
```

## Data Model

A full example of possibile options in an an Exceptionless event is below. Keep in mind that Exceptionless is a flexible API and within the `data` object, you can pass in just about any key/value pairs you want. The below model is just a reference: 

```json
{
  "type": "error",
  "source": "Website", 
  "reference_id": "123",
  "message": "some event message",
  "geo": "latitude}, longitude",
  "date":"2030-01-01T12:00:00.0000000-05:00",
  "data": {
    "@ref": {
      "id": "parent event reference id",
      "name": "parent event reference name"
    },
    "@user": {
      "identity": "email or something",
      "name": "John Doe",
      "data": "Anything we want"
    },
    "@user_description": {
      "email_address": "email",
      "description": "super cool user",
      "data": "Anything we want"
    },
    "@stack": { 
      "signature_data": {
        "ManualStackingKey": "manual key we set"
      },
      "title": "stack title"
    },
  },
  "value": "some number", 
  "tags": ["string", "string", "string"]
}
```
