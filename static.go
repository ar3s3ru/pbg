package main

import (
    "os"
    "fmt"

    "github.com/valyala/fasthttp"
)

const (
    staticPath = "/static/*resource"
    rootPath   = "/"
)

func handleRoot(ctx *fasthttp.RequestCtx) {
    fmt.Fprintf(ctx, "Hello, %s!", ctx.RemoteAddr())
}

func getStaticDirHandler() fasthttp.RequestHandler {
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    fs := &fasthttp.FS{Root: wd, GenerateIndexPages: true}
    return fs.NewRequestHandler()
}
