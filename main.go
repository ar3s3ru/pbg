package main

import (
    "os"
    "log"
    "flag"
    "errors"

    //"runtime"
    "runtime/pprof"

    "github.com/ar3s3ru/PokemonBattleGo/mem"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
    "github.com/valyala/fasthttp"
)

var (
    // Errors
    ErrInHandlerConversion = errors.New("Some error occurred, contact sysadmin, please")

    // Flags
    httpPort = flag.Int("p", 8080, "HTTP port where the server starts listening on")
    dataSet  = flag.String("f", "", "file to use as Pokèmon and Move dataset")
    cpuProf  = flag.String("cpuprofile", "", "file to use as CPU profiler dump")
    memProf  = flag.String("memprofile", "", "file to use as RAM profiler dump")
)

func withCPUProfile(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
    return func(ctx *fasthttp.RequestCtx) {
        if *cpuProf != "" {
            f, err := os.Create(*cpuProf)
            if err != nil {
                panic(err)
            }

            pprof.StartCPUProfile(f)
            defer pprof.StopCPUProfile()
        }

        handler(ctx)
    }
}

func withMEMProfile(handler fasthttp.RequestHandler) fasthttp.RequestHandler {
    return func(ctx *fasthttp.RequestCtx) {
        // Esegui prima l'handler
        handler(ctx)

        if *memProf != "" {
            f, err := os.Create(*memProf)
            if err != nil {
                panic(err)
            }

            //runtime.GC()
            pprof.WriteHeapProfile(f)
            f.Close()
        }
    }
}

func routing(server pbg.Server) {
    server.GET(StaticPath, getStaticDirHandler())
    server.GET(RootPath,   handleRoot)

    server.API_GET(PokèmonIdPath,
        pbg.Adapt(handlePokèmonId, server.WithDataAccess))
    server.API_GET(PokèmonPath,
        pbg.Adapt(handlePokèmonList, server.WithDataAccess))

    server.API_POST(RegistratonPath,
        pbg.Adapt(handleRegistration, //withMEMProfile,
                                      server.WithLogger,
                                      server.WithDataAccess))
    server.API_POST(LoginPath,
        pbg.Adapt(handleLogin, //withMEMProfile,
                               server.WithDataAccess,
                               server.WithSessionAccess))

    server.API_GET(MePath, handleMePath)
    server.API_POST(SetupPath,
        pbg.Adapt(handleSettingTeamUp, server.WithDataAccess))
}

func main() {
    flag.Parse()

    server := pbg.NewServer(
        pbg.WithHTTPPort(*httpPort),
        pbg.WithAPIResponser(pbg.NewJSONResponser()),
        pbg.WithLogger(log.New(os.Stderr, "Server - ", log.LstdFlags)),
        pbg.WithDataComponent(
            mem.NewDataComponent(
                mem.WithDatasetFile(*dataSet),
                mem.WithDataLogger(log.New(os.Stderr, "DataComponent - ", log.LstdFlags)),
            ),
        ),
        pbg.WithSessionComponent(
            mem.NewSessionComponent(
                mem.WithSessionLogger(log.New(os.Stderr, "SessionComponent - ", log.LstdFlags)),
            ),
        ),
    )

    routing(server)
    // Exits only on ListenAndServe
    server.Start()
}
