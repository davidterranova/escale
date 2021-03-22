package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	v1 "github.com/davidterranova/homework/api/protobuf/v1"
	"github.com/davidterranova/homework/client/adapters"
	"github.com/davidterranova/homework/client/repository/file"
	"github.com/davidterranova/homework/client/service"
	"google.golang.org/grpc"
)

var (
	jsonFile          = flag.String("json-ports-file", "", "the json file containing the list of ports")
	gprcServerAddress = flag.String("store-grpc-server", "", "the address of the grpc server")
)

func main() {
	flag.Parse()

	fh, err := os.Open(*jsonFile)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer fh.Close()

	conn, err := grpc.Dial(*gprcServerAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %s", err)
	}

	writer := adapters.NewPortGRPCClient(
		v1.NewPortServiceClient(conn),
	)
	reader := file.NewJSONPortLoader(fh)
	start := time.Now()
	serviceLoader := service.NewLoaderService(reader, writer)
	serviceLoader.Load()
	fmt.Printf("finished loading %d ports in %ss\n", serviceLoader.NbLoaded(), time.Since(start))
}
