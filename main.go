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

type (
    apiOkResponse struct {
        Data interface{} `json:"data"`
    }

    apiErrReponse struct {
        Error   string `json:"error"`
        Message string `json:"message"`
    }
)

const (
    CfgPokèmonFile = "POKÈMON_FILE"
    CfgLayoutFile  = "LAYOUT_FILE"

    APIEndpoint = "/api"
    StaticPath  = "/static/*resource"
)

func handleStatic(ctx *fasthttp.RequestCtx, pm fasthttprouter.Params) {
    pth := path.Join("./static", pm.ByName("resource"))
    fasthttp.ServeFile(ctx, pth)
}

func apiResponseCallback(statusCode int, data interface{}, err error, ctx *fasthttp.RequestCtx) {
    // If there is not an error, mashal and send data
    if err == nil {
        apiData := apiOkResponse{ data }
        if resp, err := json.Marshal(apiData); err == nil {
            fmt.Fprintln(ctx, string(resp))
            ctx.SetStatusCode(statusCode)
        } else {
            statusCode = fasthttp.StatusInternalServerError
        }
    }
    // If there was an error, or there wasn't an error but came up lately, send informations
    if err != nil {
        // Error response
        ne := apiErrReponse{
            Error:   err.Error(),
            Message: fasthttp.StatusMessage(statusCode),
        }
        // Mashalling error response
        if se, err := json.Marshal(ne); err != nil {
            // What a disaster...
            fmt.Fprintln(ctx, `{"error":"Internal Server Error","message": "That was a disaster..."}`)
            ctx.SetStatusCode(fasthttp.StatusInternalServerError)
        } else {
            // Phew, we did it!
            fmt.Fprintln(ctx, string(se))
            ctx.SetStatusCode(statusCode)
        }
    }
    // No matter what, API content is UTF-8 JSON
    ctx.SetContentType("application/json; charset=utf-8")
}

func newConfig(port int) pbgServer.Configuration {
    return pbgServer.NewConfig().SetHTTPPort(port).SetValue(CfgPokèmonFile, "pokedb.json")
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
    ).UseAPIResponse(apiResponseCallback).Build()
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
    wd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    fs := &fasthttp.FS{
        Root:               wd,
        GenerateIndexPages: true,
    }

    staticReqHndl := fs.NewRequestHandler()

    // Funzioni per i pokèmon
    srv.APIHandle(pbgServer.GET, APIPokèmonList, handlePokèmons)
    srv.APIHandle(pbgServer.GET, APIPokèmonEntry, handlePokèmonId)

    // Funzioni per le mosse
    srv.APIHandle(pbgServer.GET, APIMoveList, handleMoves)
    srv.APIHandle(pbgServer.GET, APIMoveEntry, handleMoveId)

    // File server!
    srv.Handle(pbgServer.GET, StaticPath,
        func (ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
            staticReqHndl(ctx)
        },
    )

    // Login e registrazione
    srv.APIHandle(pbgServer.POST, APIRegister, handleRegister)
    srv.APIHandle(pbgServer.POST, APILogin, handleLogin)
    srv.APIAuthHandle(pbgServer.GET, APIMe, handleMe)

    // Avvia il server!
    srv.StartServer()
}
