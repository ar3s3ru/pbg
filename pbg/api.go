package pbg

import (
    "encoding/json"
    "github.com/valyala/fasthttp"
)

type (
    ApiResponser interface {
        Writer(int, interface{}, interface{}) []byte
        ContentType() []byte
    }

    JsonResponser func(statusCode int, data interface{}, err interface{}) []byte

    JsonSuccess struct {
        Data interface{} `json:"data"`
    }

    JsonError struct {
        Error   string      `json:"error"`
        Message interface{} `json:"message"`
    }
)

const (
    ApiDataKey  = "apiData"
    ApiErrorKey = "apiError"

    jsonContentType = []byte("application/json; charset=utf-8")
)

func (srv server) apiWriter(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
    return func(ctx *fasthttp.RequestCtx) {
        handler(ctx)

        ctx.Response.SetBody(srv.apiResponser.Writer(
            ctx.Response.StatusCode(),
            ctx.UserValue(ApiDataKey),
            ctx.UserValue(ApiErrorKey),
        ))

        ctx.SetContentTypeBytes(srv.apiResponser.ContentType())
    }
}

func NewJsonResponser() JsonResponser {
    return func(statusCode int, data interface{}, err interface{}) []byte {
        var (
            response []byte
            errInt error
        )

        if data != nil {
            dataJson := JsonSuccess{data}
            response, errInt = json.Marshal(dataJson)
        } else if err != nil {
            errJson := JsonError{fasthttp.StatusMessage(statusCode), err}
            response, errInt = json.Marshal(errJson)
        } else {
            // TODO: insert error here (se i due casi precedenti non sono validi)
        }

        if errInt != nil {
            return []byte(`{"error":"` + fasthttp.StatusMessage(statusCode) + `","message":"` + errInt.Error() + `"}`)
        }

        return response
    }
}

func (jr JsonResponser) Writer(statusCode int, data interface{}, err interface{}) []byte {
    return jr(statusCode, data, err)
}

func (jr JsonResponser) ContentType() []byte {
    return jsonContentType
}
