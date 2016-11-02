package main

import (
    "github.com/ar3s3ru/PokemonBattleGo/mem"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

func createDataMechanism() pbg.DataMechanism {
    db := mem.NewDataBuilder().WithSourceFile("pokedb.json")
    return db.Build()
}

func createServer(cfg pbg.Configuration,
                  dm pbg.DataMechanism, sm pbg.SessionMechanism, am pbg.AuthorizationMechanism) pbg.Server {
    // Server builder
    srvBuild := pbg.NewServerBuilder().
                WithConfiguration(cfg).
                WithDataMechanism(dm).
                WithSessionMechanism(sm).
                WithAuthorizationMechanism(am)

    return srvBuild.Build()
}

func main() {
    config := pbg.BaseConfiguration{}

    dataMechanism := createDataMechanism()
    sessionMechanism := mem.NewSessionMechanism()
    authorizationMechanism, _ := sessionMechanism.(pbg.AuthorizationMechanism)

    server := createServer(
        config, dataMechanism, sessionMechanism, authorizationMechanism,
    )

    // Start server loop
    server.Start()
}
