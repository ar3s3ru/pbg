package main

import (
    "fmt"

    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "github.com/valyala/fasthttp"
)

type unitTest func(pbg.Server, *fasthttp.HostClient)

var (
    client *fasthttp.HostClient
    config *pbg.BaseConfiguration

    apiAddr string
)

func init() {

    if server != nil {
        panic("Server should be nil here")
    }

    config  = &pbg.BaseConfiguration{Port: 8080, Local: true, Endpoint: "/api"}
    apiAddr = fmt.Sprintf("http://localhost:%d%s", config.Port, config.Endpoint)

    // Client thing
    client = &fasthttp.HostClient{
        Addr: fmt.Sprintf("localhost:%d", config.Port),
    }

    // Server thing
    sessionMechanism := createSessionMechanism()
    createServer(
        config,
        createDataMechanism(),
        sessionMechanism,
        createAuthorizationMechanism(sessionMechanism),
    )

    // Handle some things here
    setupServer()
    startAsynchronous()
}

func withServerTesting(test unitTest) {
    test(server, client)
}

