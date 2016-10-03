package main

import (
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer/mem"
    "path"
    "fmt"
    "encoding/json"
)

const (
    CfgPokèmonFile = "POKÈMON_FILE"
    CfgLayoutFile  = "LAYOUT_FILE"
)

func handleStatic(_ pbgServer.IServerContext, ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    pth := path.Join("static", pm.ByName("resource"))
    fasthttp.ServeFile(ctx, pth)
}

func getServer() pbgServer.PBGServer {
    return pbgServer.Builder(
        // Configura il server
        pbgServer.NewConfig().SetHTTPPort(8080).SetValue(
            CfgPokèmonFile, "pokedb.json",
        ).SetValue(
            CfgLayoutFile, "templates/layout.html",
        ).SetValue(
            CfgPokemonListFile, "templates/pokemons.html",
        ).SetValue(
            CfgPokemonIdFile, "templates/pokemon_id.html",
        ).SetValue(
            CfgMovesListFile, "templates/moves.html",
        ).SetValue(
            CfgMovesIdFile, "templates/move_id.html",
        ),
    ).UseDataMechanism(
        // Data mechanism callback
        func (cfg pbgServer.Configuration) pbgServer.IDataMechanism {
            if pokèmonFile := cfg.GetValue(CfgPokèmonFile); pokèmonFile == nil {
                panic("PokèmonFile not configured")
            } else {
                return mem.NewDataBuilder().UsePokèmonFile(pokèmonFile.(string)).Build()
            }
        },
    ).UseSessMechanism(
        // Session mechanism callback
        func (_ pbgServer.Configuration, dm pbgServer.IDataMechanism) pbgServer.ISessionMechanism {
            // Crea un nuovo oggetto Authority dal package mem
            return mem.AuthBuilder().UseDataMechanism(dm).Build()
        },
    ).UseAuthMechanism(
        func (_ pbgServer.Configuration, sm pbgServer.ISessionMechanism) pbgServer.IAuthMechanism {
            // SM è di tipo Authority, quindi usalo come AuthMechanism
            return sm.(mem.IAuthority)
        },
    ).Build()
}

func main() {
    // Get server instance
    srv := getServer()

    // Funzioni per i pokèmon
    srv.Handle(
        pbgServer.GET, "/pokemons", handlePokèmons,
    ).Handle(
        pbgServer.GET, "/pokemon/:id", handlePokèmonId,
    )
    // Funzioni per le mosse
    srv.Handle(
        pbgServer.GET, "/moves", handleMoves,
    ).Handle(
        pbgServer.GET, "/move/:id", handleMoveId,
    )
    // File server!
    srv.Handle(pbgServer.GET, "/static/*resource", handleStatic)
    // Login e registrazione
    srv.Handle(
        pbgServer.POST, "/register", handleRegister,
    ).Handle(
        pbgServer.GET, "/register", handleGetRegister,
    ).Handle(
        pbgServer.POST, "/login", handleLogin,
    ).AuthHandle(
        pbgServer.GET, "/me",
        func(sess pbgServer.Session, sctx pbgServer.IServerContext, ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
            lel, _ := json.Marshal(sess.GetUserReference())
            fmt.Fprintf(ctx, "It's me, %s\n", lel)
        },
    )
    // Avvia il server!
    srv.StartServer()
}
