package main

import (
    "github.com/valyala/fasthttp"
    "github.com/buaazp/fasthttprouter"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer"
    "github.com/ar3s3ru/PokemonBattleGo/pbgServer/mem"
    "path"
    "fmt"
    "encoding/json"
    "strconv"
    "os"
)

const (
    CfgPokèmonFile = "POKÈMON_FILE"
    CfgLayoutFile  = "LAYOUT_FILE"

    APIEndpoint = "/api"
    StaticPath  = "/static"
)

func handleStatic(ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    pth := path.Join("static", pm.ByName("resource"))
    fasthttp.ServeFile(ctx, pth)
}

func newConfig(port int) pbgServer.Configuration {
    return pbgServer.NewConfig().SetHTTPPort(port).SetValue(
        CfgPokèmonFile, "pokedb.json",
    )//.SetValue(
    //    CfgLayoutFile, path.Join("templates", "layout.html"),
    //).SetValue(
    //    CfgPokemonListFile, path.Join("templates", "pokemons.html"),
    //).SetValue(
    //    CfgPokemonIdFile, path.Join("templates", "pokemon_id.html"),
    //).SetValue(
    //    CfgMovesListFile, path.Join("templates", "moves.html"),
    //).SetValue(
    //    CfgMovesIdFile, path.Join("templates", "move_id.html"),
    //)
}

func getServer(port int) pbgServer.PBGServer {
    return pbgServer.Builder(newConfig(port)).UseDataMechanism(
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
    ).UseAPIResponse(
        func (statusCode int, data interface{}, err error, ctx *fasthttp.RequestCtx) {
            type (
                d struct {
                    Data interface{} `json:"data"`
                }

                e struct {
                    Error   string `json:"error"`
                    Message string `json:"message"`
                }
            )


            if err == nil {
                nd := d{data}
                if resp, err := json.Marshal(nd); err == nil {
                    fmt.Fprintln(ctx, string(resp))
                    ctx.SetStatusCode(statusCode)
                } else {
                    statusCode = fasthttp.StatusInternalServerError
                }
            }

            if err != nil {
                ne := e{fasthttp.StatusMessage(statusCode), err.Error()}
                if se, err := json.Marshal(ne); err != nil {
                    fmt.Fprintln(ctx, "{\"error\":\"Internal Server Error\",\"message\": \"That was a disaster...\"}")
                    ctx.SetStatusCode(fasthttp.StatusInternalServerError)
                } else {
                    fmt.Fprintln(ctx, string(se))
                    ctx.SetStatusCode(statusCode)
                }
            }

            ctx.SetContentType("application/json; charset=utf-8")
        },
    ).Build()
}

func main() {
    var port = 8080
    if len(os.Args) >= 2 {
        // Get server instance
        p, err := strconv.Atoi(os.Args[1])
        if err == nil {
            port = p
        }
    }

    srv := getServer(port)

    // Funzioni per i pokèmon
    srv.APIHandle(pbgServer.GET, APIPokèmonList, handlePokèmons)
    srv.APIHandle(pbgServer.GET, APIPokèmonEntry, handlePokèmonId)

    // Funzioni per le mosse
    srv.APIHandle(pbgServer.GET, "/moves", handleMoves)
    srv.APIHandle(pbgServer.GET, "/move/:id", handleMoveId)

    // File server!
    srv.Handle(pbgServer.GET, StaticPath, handleStatic)

    // Login e registrazione
    srv.APIHandle(
        pbgServer.POST, "/register", handleRegister,
    ).APIHandle(
        pbgServer.POST, "/login", handleLogin,
    ).APIAuthHandle(
        pbgServer.GET, "/me",
        // JSON handler
        func (sess pbgServer.Session, _ pbgServer.IServerContext,
               ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) (int, interface{}, error) {
            return fasthttp.StatusOK, sess.GetUserReference(), nil
        },
    )
    // Avvia il server!
    srv.StartServer()
}
