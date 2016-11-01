package main

import (
    "github.com/ar3s3ru/PokemonBattleGo/mem"
    "github.com/ar3s3ru/PokemonBattleGo/pbg"
)

func createDataMechanism() pbg.DataMechanism {
    db := mem.NewDataBuilder().WithSourceFile("pokedb.json")
    return db.Build()
}

func createServer(cfg pbg.Configuration, dm pbg.DataMechanism) pbg.Server {
    srvBuild := pbg.NewServerBuilder().WithConfiguration(cfg).WithDataMechanism(dm)
    return srvBuild.Build()
}

func main() {
    config := pbg.BaseConfiguration{

    }

    dataMechanism := createDataMechanism()
    //sessionMechanism := ""
    //authorizationMechanism := ""

    server := createServer(config, dataMechanism)

    server.Start()
}
