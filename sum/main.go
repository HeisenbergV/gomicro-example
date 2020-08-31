package main

import (
	"context"
	"fmt"
	"gomicro-example/proto/numone"
	"gomicro-example/proto/numtwo"

	"github.com/micro/go-micro/v2"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.sum"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	one := numone.NewNumService("go.micro.srv.numone", service.Client())
	oneResponse, err := one.GetNum(context.TODO(), &numone.Request{})
	fmt.Println(err)
	two := numtwo.NewNumService("go.micro.srv.numtwo", service.Client())
	twoResponse, _ := two.GetNum(context.TODO(), &numtwo.Request{})

	sum := oneResponse.Num + twoResponse.Num
	fmt.Println(sum)
}
