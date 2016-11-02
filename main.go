package main

import (
    "github.com/ar3s3ru/PokemonBattleGo/mem"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

func createDataMechanism() pbg.DataMechanism {
    db := mem.NewDataBuilder().WithSourceFile("pokedb.json")
    return db.Build()
}

func createServer(dm pbg.DataMechanism, sm pbg.SessionMechanism, am pbg.AuthorizationMechanism) pbg.Server {
    // Server builder
    srvBuild := pbg.NewServerBuilder().
                WithDataMechanism(dm).
                WithSessionMechanism(sm).
                WithAuthorizationMechanism(am)

    return srvBuild.Build()
}

func main() {
    dataMechanism          := createDataMechanism()
    sessionMechanism       := mem.NewSessionMechanism()
    authorizationMechanism := sessionMechanism.(pbg.AuthorizationMechanism)

    server := createServer(dataMechanism, sessionMechanism, authorizationMechanism)

    // Handle some shit here
    server.Handle(pbg.GET, staticPath, handleStaticPath())
    server.Handle(pbg.GET, rootPath, handleRoot)

    server.APIHandle(pbg.GET, pokèmonIdPath,
        pbg.Adapt(handlePokèmonId, server.WithDataAccess))
    server.APIHandle(pbg.GET, pokèmonPath,
        pbg.Adapt(handlePokèmonList, server.WithDataAccess))

    // Start server loop
    server.Start()
}
