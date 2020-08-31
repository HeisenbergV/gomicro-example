package subscriber

import (
	"context"

	num "gomicro-example/proto/numtwo"
)

type Num struct{}

func (e *Num) Handle(ctx context.Context, msg *num.Message) error {
	return nil
}

func Handler(ctx context.Context, msg *num.Message) error {
	return nil
}
