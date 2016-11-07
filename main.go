package main

import (
    "errors"

    "github.com/ar3s3ru/PokemonBattleGo/mem"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

var (
    // Server
    server pbg.Server
    // Errors
    ErrInHandlerConversion = errors.New("Some error occurred, contact sysadmin, please")
)

func createDataMechanism() pbg.DataMechanism {
    db := mem.NewDataBuilder().WithSourceFile("pokedb.json")
    return db.Build()
}

func createSessionMechanism() pbg.SessionMechanism {
    return mem.NewSessionMechanism()
}

func createAuthorizationMechanism(sm pbg.SessionMechanism) pbg.AuthorizationMechanism {
    if am, ok := sm.(pbg.AuthorizationMechanism); !ok {
        panic("Invalid SessionMechanism used here")
    } else {
        return am
    }
}

func createServer(cfg pbg.Configuration, dm pbg.DataMechanism, sm pbg.SessionMechanism, am pbg.AuthorizationMechanism) {
    // Server builder
    srvBuild := pbg.NewServerBuilder().
                WithDataMechanism(dm).
                WithSessionMechanism(sm).
                WithAuthorizationMechanism(am)

    if cfg != nil {
        srvBuild.WithConfiguration(cfg)
    }

    server = srvBuild.Build()
}

func setupServer() {
    server.Handle(pbg.GET, StaticPath, getStaticDirHandler())
    server.Handle(pbg.GET, RootPath, handleRoot)

    server.APIHandle(pbg.GET, PokèmonIdPath,
        pbg.Adapt(handlePokèmonId, server.WithDataAccess))
    server.APIHandle(pbg.GET, PokèmonPath,
        pbg.Adapt(handlePokèmonList, server.WithDataAccess))

    server.APIHandle(pbg.POST, RegistratonPath,
        pbg.Adapt(handleRegistration, server.WithDataAccess))
    server.APIHandle(pbg.POST, LoginPath,
        pbg.Adapt(handleLogin, server.WithDataAccess, server.WithSessionAccess))

    server.APIAuthHandle(pbg.GET, MePath, handleMePath)
    server.APIAuthHandle(pbg.POST, SetupPath,
        pbg.Adapt(handleSettingTeamUp, server.WithDataAccess))
}

func startSynchronous() {
    server.Start()
}

func startAsynchronous() {
    go server.Start()
}

func main() {

    if server != nil {
        panic("server already started")
    }

    sessionMechanism := createSessionMechanism()
    createServer(
        nil,
        createDataMechanism(),
        sessionMechanism,
        createAuthorizationMechanism(sessionMechanism),
    )

    // Handle some things here
    setupServer()

    // Start server loop
    startSynchronous()
}
