// source: num.proto

package numone

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"

	context "context"

	client "github.com/micro/go-micro/v2/client"

	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Num service

type NumService interface {
	GetNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type numService struct {
	c    client.Client
	name string
}

func NewNumService(name string, c client.Client) NumService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "numone"
	}
	return &numService{
		c:    c,
		name: name,
	}
}

func (c *numService) GetNum(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Num.GetNum", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Num service

type NumHandler interface {
	GetNum(context.Context, *Request, *Response) error
}

func RegisterNumHandler(s server.Server, hdlr NumHandler, opts ...server.HandlerOption) error {
	type num interface {
		GetNum(ctx context.Context, in *Request, out *Response) error
	}
	type Num struct {
		num
	}
	h := &numHandler{hdlr}
	return s.Handle(s.NewHandler(&Num{h}, opts...))
}

type numHandler struct {
	NumHandler
}

func (h *numHandler) GetNum(ctx context.Context, in *Request, out *Response) error {
	return h.NumHandler.GetNum(ctx, in, out)
}
