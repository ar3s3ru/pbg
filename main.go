package main

import (
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "fmt"
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
)

func main() {
    srv := pbgServer.Builder().UseConfiguration(
        pbgServer.ObjectConfigurationBuilder(8080),
    ).UseDataMechanism(func(cfg pbgServer.IConfiguration) pbgServer.IDataMechanism {
        fmt.Printf("Building Data Mechanism using %v\n", cfg)
        return nil
    }).UseAuthMechanism(func (cfg pbgServer.IConfiguration) pbgServer.IAuthMechanism {
        fmt.Printf("Building Auth Mechanism using %v\n", cfg)
        return nil
    }).UseSessMechanism(func (cfg pbgServer.IConfiguration) pbgServer.ISessionsMechanism {
        fmt.Printf("Building Sess Mechanism using %v\n", cfg)
        return nil
    }).Build()

    srv.Handle(pbgServer.GET, "/hello", func(sctx pbgServer.IServerContext,
                                             ctx *fasthttp.RequestCtx,
                                             pm fasthttprouter.Params) {
        fmt.Fprintf(ctx, "Called \"/hello\" with %v\n", sctx)
    })

    srv.StartServer()
}
