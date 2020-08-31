package handler

import (
	"context"

	num "gomicro-example/proto/num-two"
)

type Num struct{}

func (e *Num) Call(ctx context.Context, req *num.Request, rsp *num.Response) error {
	return nil
}
