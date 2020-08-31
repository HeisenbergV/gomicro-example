package main

import (
	"gomicro-example/num-one/handler"
	"gomicro-example/proto/numone"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.numone"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	numone.RegisterNumHandler(service.Server(), new(handler.Num))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
