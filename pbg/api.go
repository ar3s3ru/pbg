package pbg

import (
    "encoding/json"
    "github.com/valyala/fasthttp"
)

type (
    APIResponser interface {
        Writer(int, interface{}, interface{}) []byte
        ContentType() []byte
    }

    JSONResponser func(statusCode int, data interface{}, err interface{}) []byte

    JSONSuccess struct {
        Data interface{} `json:"data"`
    }

    JSONError struct {
        Error   string      `json:"error"`
        Message interface{} `json:"message"`
    }
)

const (
    APIDataKey  = "apiData"
    APIErrorKey = "apiError"
)

var jsonContentType = []byte("application/json; charset=utf-8")

func (srv *server) apiWriter(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
    return func(ctx *fasthttp.RequestCtx) {
        handler(ctx)

        ctx.Response.SetBody(srv.apiResponser.Writer(
            ctx.Response.StatusCode(),
            ctx.UserValue(APIDataKey),
            ctx.UserValue(APIErrorKey),
        ))

        ctx.SetContentTypeBytes(srv.apiResponser.ContentType())
    }
}

func NewJSONResponser() JSONResponser {
    return func(statusCode int, data interface{}, err interface{}) []byte {
        var (
            response []byte
            errInt error
        )

        if data != nil {
            dataJson := JSONSuccess{data}
            response, errInt = json.Marshal(dataJson)
        } else if err != nil {
            errJson := JSONError{fasthttp.StatusMessage(statusCode), err}
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

func (jr JSONResponser) Writer(statusCode int, data interface{}, err interface{}) []byte {
    return jr(statusCode, data, err)
}

func (jr JSONResponser) ContentType() []byte {
    return jsonContentType
}
