package life

import (
	"context"
	"time"
)

// CTXTimeout ...
func CTXTimeout() (context.Context, context.CancelFunc) {
	return context.WithTimeout(CTX, 30*time.Second)
}
