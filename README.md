# Exceptionless Go Client

This is a simple client wrapper to make working with Exceptionless in the Go environment easier. The client, is not a complete wrapper around the Exceptionless API, so you are encouraged to review the [full Swagger documentation here](https://api.exceptionless.io/docs/index.html). 

## Getting Started 

To install the client, run `go get https://github.com/Exceptionless/go-client`.
 Import it into your project like this: 

```go
import (
	"github.com/Exceptionless/go-client"
)
```

Once you've imported it in your project, you'll need to configure the client. The current configuration options are as follows: 

* apiKey - string - required
* serverURL - string - optional
* updateSettingsWhenIdleInterval - int32 - optional
* includePrivateInformation - bool - optional

`apiKey` is self-explanatory. Get your key at https://exceptionless.com.

If you are self-hosting Exceptionless, provide the server URL for your self-hosted installation for the `serverURL` property. 

If you would like to update configuration settings within the Exceptionless frontend and allow the Go Client to adjust itself based on those settings, you can set a polling interval for `updateSettingsWhenIdleInterval`. If this is configured, the client will automatically poll for your project configuration settings. 

If you pass in a `false` for `includePrivateInformation`, the Go Client will try to strip out information like passwords, api keys, bearer tokens, etc. 

## Sending Events  

This client has two convenience functions: `SubmitError` and `SubmitLog`. These functions take minimal arguments and are far less flexible than building events yourself, but they are the easiest to use. We'll cover them first. 

### SubmitError

