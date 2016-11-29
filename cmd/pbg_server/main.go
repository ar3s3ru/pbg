package main

import (
	"errors"
	"flag"
	"log"
	"os"

	"time"
	//"runtime"
	//"syscall"
	//"os/signal"
	"runtime/pprof"

	"github.com/ar3s3ru/pbg"
	"github.com/ar3s3ru/pbg/mem"
)

var (
	// Errors
	ErrInHandlerConversion = errors.New("Some error occurred, contact sysadmin, please")

	// Flags
	httpPort = flag.Int("p", 8080, "HTTP port where the server starts listening on")
	dataSet  = flag.String("f", "pokedb.json", "file to use as Pokèmon and Move dataset")
	cpuProf  = flag.String("cpuprofile", "", "file to use as CPU profiler dump")
	memProf  = flag.String("memprofile", "", "file to use as RAM profiler dump")
)

func CPUSampling() {
	if *cpuProf != "" {
		f, err := os.Create(*cpuProf)
		if err != nil {
			panic(err)
		}

		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()

		time.Sleep(2 * time.Second)
		log.Println("CPU profile terminated")
	}
}

func MEMSampling() {
	if *memProf != "" {
		f, err := os.Create(*memProf)
		if err != nil {
			panic(err)
		}

		//runtime.GC()
		pprof.WriteHeapProfile(f)
		f.Close()

		log.Println("MEM profile terminated")
	}
}

func notifyHandling(sigCh <-chan os.Signal) {
	//for {
	//    sig := <-sigCh
	//    if sig == syscall.SIGUSR1 {
	//        log.Println("Received SIGUSR1")
	//        go CPUSampling()
	//    } else if sig == syscall.SIGUSR2 {
	//        log.Println("Received SIGUSR2")
	//        go MEMSampling()
	//    }
	//}
}

func registerSignals() {
	sigCh := make(chan os.Signal, 1)
	//signal.Notify(sigCh, syscall.SIGUSR1, syscall.SIGUSR2)

	go notifyHandling(sigCh)
}

func creating(moves []pbg.Move, pokèmons []pbg.Pokèmon) pbg.Server {
	return pbg.NewServer(
		pbg.WithHTTPPort(*httpPort),
		pbg.WithAPIResponser(pbg.NewJSONResponser()),
		pbg.WithLogger(log.New(os.Stderr, "Server - ", log.LstdFlags|log.Lshortfile)),
		pbg.WithMoveDBComponent(
			mem.NewMoveDBComponent(
				mem.WithMoves(moves),
				mem.WithMoveDBLogger(log.New(os.Stderr, "MoveDB - ", log.LstdFlags|log.Lshortfile)),
			),
		),
		pbg.WithPokèmonDBComponent(
			mem.NewPokèmonDBComponent(
				mem.WithPokèmons(pokèmons),
				mem.WithPokèmonDBLogger(log.New(os.Stderr, "PokèmonDB - ", log.LstdFlags|log.Lshortfile)),
			),
		),
		pbg.WithTrainerDBComponent(
			mem.NewTrainerDBComponent(
				mem.WithTrainerDBLogger(log.New(os.Stderr, "TrainerDB - ", log.LstdFlags|log.Lshortfile)),
			),
		),
		pbg.WithSessionDBComponent(
			mem.NewSessionComponent(
				mem.WithSessionDBLogger(log.New(os.Stderr, "SessionDB - ", log.LstdFlags|log.Lshortfile)),
			),
		),
	)
}

func routing(server pbg.Server) {
	withAuthorization := authorizedHandler(server)

	server.GET(StaticPath, getStaticDirHandler())
	server.GET(RootPath, handleRoot)

	server.API_GET(PokèmonIdPath, pbg.Adapt(handlePokèmonId, server.WithPokèmonDBAccess))
	server.API_GET(PokèmonPath, pbg.Adapt(handlePokèmonList, server.WithPokèmonDBAccess))

	server.API_POST(RegistratonPath, pbg.Adapt(HandleRegistration, server.WithTrainerDBAccess))
	server.API_POST(LoginPath, pbg.Adapt(HandleLogin, server.WithTrainerDBAccess, server.WithSessionDBAccess))

	server.API_GET(MePath, pbg.Adapt(handleMePath, withAuthorization))
	server.API_POST(SetupPath, pbg.Adapt(handleSettingTeamUp,
			withAuthorization, server.WithMoveDBAccess, server.WithPokèmonDBAccess))
}

func main() {
	flag.Parse()

	pDataset, mDataset, err := mem.WithDatasetFile(*dataSet)
	if err != nil {
		log.Fatalln(err)
	}

	// Create new server
	server := creating(mDataset, pDataset)

	// Register handlers and signals notifier
	routing(server)
	registerSignals()

	// Exits only on ListenAndServe
	server.Start()
}
