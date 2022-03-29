package u2utils

import (
	"context"
	"time"
)

func Ctx(second time.Duration) context.Context {
	ctx, _ := context.WithTimeout(context.Background(), second*time.Second)
	return ctx
}
