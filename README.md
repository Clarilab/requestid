# requestid

## Description

this repository is a collection of middlewares, used to get a requestID from an http request, and set it in rest clients used to perform requests.

## supported frameworks

| Framework | Support | Version |
|:----------|:-------:|:--------:|
| Resty     |    ✓    |      v2 |
| Atreugo   |    ✓    |     v11 |

## how to install

```bash
go get github.com/Clarilab/requestid@v0.2.0
```

## How to use

### resty

```go
package main

import (
	"context"

	"github.com/Clarilab/requestid"
	"github.com/go-resty/resty/v2"
)

func main(){
	client := resty.
		New().
		OnBeforeRequest(requestid.RestyMiddleware)

	// this most likely will be set by an http framework such as atreugo.
	ctx := context.TODO()
	ctx = requestid.Set(ctx, "<TheRequestID>")

	resp, err := client.R().
		SetContext(ctx).
		Get("https://example.com")

	if err != nil {
		panic(err)
	}

	println(string(resp.Body()))
}

```

### atreugo

```go
package main

import (
	"github.com/Clarilab/requestid"
	atreugoRid "github.com/atreugo/requestid"
	"github.com/savsgio/atreugo/v11"
	"github.com/valyala/fasthttp"
	"log"
)

func main() {
	router := atreugo.New(atreugo.Config{Addr: "0.0.0.0:1337"})
	router.UseBefore(atreugoRid.New(atreugoRid.Config{}))
	router.UseBefore(requestid.AtreugoMiddleware())

	router.GET("/", func(c *atreugo.RequestCtx) error {
		rid, err := requestid.Get(c.AttachedContext())
		if err != nil {
			println(err.Error())

			c.Response.SetStatusCode(fasthttp.StatusNoContent)

			return nil
		}

		println(rid)

		c.Response.SetStatusCode(fasthttp.StatusNoContent)

		return nil
	})

	log.Fatal(router.ListenAndServe())
}
```
