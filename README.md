# Exceptionless Go Client

This is a simple client wrapper to make working with Exceptionless in the Go environment easier. The client, is not a complete wrapper around the Exceptionless API, so you are encouraged to review the [full Swagger documentation here](https://api.exceptionless.io/docs/index.html). 

## Getting Started 

To install the client, run `go get https://github.com/Exceptionless/go-client`. Import it into your project like this: 

```go
import (
	"github.com/Exceptionless/go-client"
)
```

Once you've imported it in your project, you'll need to configure the client. The current configuration options are as follows: 

* apiKey - required
* 