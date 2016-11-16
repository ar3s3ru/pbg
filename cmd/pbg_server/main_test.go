package main

import (
	"log"

	"github.com/valyala/fasthttp"

	"github.com/ar3s3ru/pbg"
	"github.com/ar3s3ru/pbg/mem"
)

type (
	unitTest func(pbg.Server, *fasthttp.HostClient)
)

var (
	unitTestChannel = make(chan unitTest)
)

const (
	address     = "localhost:8080"
	httpAddress = "http://" + address + "/api"
)

func init() {
	log.Println("Initing main_test.go")
	go unitTestingLoop(unitTestChannel)
}

func creatingTestingServer() pbg.Server {
	return pbg.NewServer(
		pbg.WithHTTPPort(8080),
		pbg.WithAPIResponser(pbg.NewJSONResponser()),
		pbg.WithMoveDBComponent(
			mem.NewMoveDBComponent(
				mem.WithMoves(make([]pbg.Move, 10)),
			),
		),
		pbg.WithPokèmonDBComponent(
			mem.NewPokèmonDBComponent(
				mem.WithPokèmons(make([]pbg.Pokèmon, 10)),
			),
		),
		pbg.WithTrainerDBComponent(
			mem.NewTrainerDBComponent(),
		),
		pbg.WithSessionDBComponent(
			mem.NewSessionComponent(),
		),
	)
}

func unitTestingLoop(testRequests <-chan unitTest) {
	server := creatingTestingServer()
	client := &fasthttp.HostClient{Addr: address}

	routing(server)
	go server.Start()

	for request := range testRequests {
		request(server, client)
	}
}

func withServerTesting(test unitTest) {
	sync := make(chan bool, 1)

	unitTestChannel <- func(server pbg.Server, client *fasthttp.HostClient) {
		test(server, client)
		sync <- true
	}

	<-sync
}
