package handler

import (
	"context"

	num "gomicro-example/proto/numone"
)

type Num struct{}

func (e *Num) GetNum(ctx context.Context, req *num.Request, rsp *num.Response) error {
	rsp.Num = 1
	return nil
}
