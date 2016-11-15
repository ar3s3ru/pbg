package main

import (
    "testing"

    "github.com/valyala/fasthttp"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

type (
    signingTest struct {
        postBody *PostBody
        result   int
    }
)

var (
    registrationTests = [...]signingTest{
        {&PostBody{"dan", "lol"}, fasthttp.StatusCreated},
        {&PostBody{"dan", "lol"}, fasthttp.StatusBadRequest},
    }

    loginTests = [...]signingTest{
        {&PostBody{"dan", "lol"}, fasthttp.StatusCreated},
        {&PostBody{"danA", "lol"}, fasthttp.StatusNotFound},
    }
)

func TestHandleRegistration(t *testing.T) {
    withServerTesting(func(srv pbg.Server, clt *fasthttp.HostClient) {
        for _, test := range registrationTests {
            var (
                request fasthttp.Request
                response fasthttp.Response
            )

            body, err := EncodePostBody(test.postBody)
            if err != nil {
                t.Fatal(err.Error())
            }

            request.Header.SetMethod("POST")
            request.SetRequestURI(httpAddress + RegistratonPath)
            request.SetBody(body)

            clt.Do(&request, &response)

            if response.StatusCode() != test.result {
                t.Errorf("Status code mismatch, check results:\n\tBody: %s\n\tCode: %d (should be %d)\n",
                    response.Body(), response.StatusCode(), test.result,
                )
            }
        }
    })
}

func TestHandleLogin(t *testing.T) {
    withServerTesting(func(srv pbg.Server, clt *fasthttp.HostClient) {
        for _, test := range loginTests {
            var (
                request fasthttp.Request
                response fasthttp.Response
            )

            body, err := EncodePostBody(test.postBody)
            if err != nil {
                t.Fatal(err.Error())
            }

            request.Header.SetMethod("POST")
            request.SetRequestURI(httpAddress + LoginPath)
            request.SetBody(body)

            clt.Do(&request, &response)

            if response.StatusCode() != test.result {
                t.Errorf("Status code mismatch, check results:\n\tBody: %s\n\tCode: %d (should be %d)\n",
                    response.Body(), response.StatusCode(), test.result,
                )
            }
        }
    })
}
