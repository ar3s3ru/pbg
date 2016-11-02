package main

import (
    "github.com/valyala/fasthttp"
    "os"
    "fmt"
)

const (
    staticPath = "/static/*resource"
    rootPath   = "/"
)

func handleRoot(ctx *fasthttp.RequestCtx) {
    fmt.Fprintf(ctx, "Hello, %s!", ctx.RemoteAddr())
}

func handleStaticPath() fasthttp.RequestHandler {
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    fs := &fasthttp.FS{Root: wd, GenerateIndexPages: true}
    return fs.NewRequestHandler()
}
