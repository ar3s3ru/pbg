package pbg

import (
    "encoding/json"
    "github.com/valyala/fasthttp"
)

type (
    APIResponser interface {
        Writer(int, interface{}, error) []byte
        ContentType() []byte
    }

    JSONResponser func(statusCode int, data interface{}, err error) []byte

    JSONSuccess struct {
        Data interface{} `json:"data"`
    }

    JSONError struct {
        Error   string `json:"error"`
        Message string `json:"message"`
    }
)

var (
    jsonContentType = []byte("application/json; charset=utf-8")
    typeAssertError = []byte(`{"error":"Internal Server Error","message":"Some error occurred"}`)
)

func (srv *server) apiWriter(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
    return func(ctx *fasthttp.RequestCtx) {
        handler(ctx)

        var err error = nil
        if apiErr := ctx.UserValue(APIErrorKey); apiErr != nil {
            var ok bool
            if err, ok = apiErr.(error); !ok {
                ctx.SetStatusCode(fasthttp.StatusInternalServerError)
                ctx.SetBody(typeAssertError)
                return
            }
        }

        ctx.Response.SetBody(srv.apiResponser.Writer(
            ctx.Response.StatusCode(),
            ctx.UserValue(APIDataKey),
            err,
        ))

        ctx.SetContentTypeBytes(srv.apiResponser.ContentType())
    }
}

func NewJSONResponser() JSONResponser {
    return func(statusCode int, data interface{}, err error) []byte {
        var (
            response []byte
            errInt error
        )

        if data != nil {
            dataJson := JSONSuccess{data}
            response, errInt = json.Marshal(dataJson)
        } else if err != nil {
            errJson := JSONError{fasthttp.StatusMessage(statusCode), err.Error()}
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

func (jr JSONResponser) Writer(statusCode int, data interface{}, err error) []byte {
    return jr(statusCode, data, err)
}

func (jr JSONResponser) ContentType() []byte {
    return jsonContentType
}

func writeAPI(ctx *fasthttp.RequestCtx, key string, ent interface{}, stat int) {
    ctx.SetUserValue(key, ent)
    ctx.SetStatusCode(stat)
}

func WriteAPISuccess(ctx *fasthttp.RequestCtx, data interface{}, statusCode int) {
    writeAPI(ctx, APIDataKey, data, statusCode)
}

func WriteAPIError(ctx *fasthttp.RequestCtx, err error, statusCode int) {
    writeAPI(ctx, APIErrorKey, err, statusCode)
}
