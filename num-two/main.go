package main

import (
	"gomicro-example/num-two/handler"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	num "gomicro-example/proto/numtwo"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.numtwo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	num.RegisterNumHandler(service.Server(), new(handler.Num))
	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
