package main

import (
	"fmt"
	"os"

	"io/ioutil"
	"path/filepath"

	"github.com/valyala/fasthttp"
)

const (
	StaticPath = "/static/*resource"
	RootPath   = "/"
)

func handleRoot(ctx *fasthttp.RequestCtx) {
	if file, err := ioutil.ReadFile(filepath.Join("templates", "layout.html")); err != nil {
		ctx.Error(err.Error(), fasthttp.StatusInternalServerError)
	} else {
		fmt.Fprintf(ctx, "%s", file)
		ctx.SetContentType("text/html")
	}
}

func getStaticDirHandler() fasthttp.RequestHandler {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fs := &fasthttp.FS{Root: wd, GenerateIndexPages: true}
	return fs.NewRequestHandler()
}
