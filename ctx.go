package life

import (
	"context"
	"time"
)

// CTXTimeout 给出相应 timeout context，如果有信号则提前退出
func CTXTimeout(d time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(CTX, d)
}
