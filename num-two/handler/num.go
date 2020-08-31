package handler

import (
	"context"

	num "gomicro-example/proto/numtwo"
)

type Num struct{}

func (e *Num) GetNum(ctx context.Context, req *num.Request, rsp *num.Response) error {
	rsp.Num = 2
	return nil
}
