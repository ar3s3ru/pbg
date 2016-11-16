package main

import (
	"fmt"
	"os"

	"github.com/valyala/fasthttp"
)

const (
	StaticPath = "/static/*resource"
	RootPath   = "/"
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
