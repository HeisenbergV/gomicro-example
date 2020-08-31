package main

import (
	"gomicro-example/num-two/handler"
	"gomicro-example/num-two/subscriber"

	"github.com/micro/v2/go-micro"
	"github.com/micro/go-micro/v2/util/log"

	num "gomicro-example/proto/num-two/num"
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
	num.RegisterNumHandler(service.Server(), new(handler.Num))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.num", service.Server(), new(subscriber.Num))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.num", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
