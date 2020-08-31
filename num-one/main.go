package main

import (
	"gomicro-exmple/num-one/handler"
	"gomicro-exmple/num-one/subscriber"
	"gomicro-exmple/proto/numone"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.num"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	numone.RegisterNumHandler(service.Server(), new(handler.Num))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.num", service.Server(), new(subscriber.Num))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.num", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
