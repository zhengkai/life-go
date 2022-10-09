package life

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"syscall"
)

// ...
var (
	Stop         bool
	CTX          context.Context
	cancel       context.CancelFunc
	BeforeCancel func()

	ErrStop = errors.New(`life stop`)
)

func init() {

	CTX, cancel = context.WithCancel(context.Background())

	go func() {

		c := make(chan os.Signal, 1)
		signal.Notify(
			c,
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
		<-c
		Stop = true
		if BeforeCancel != nil {
			BeforeCancel()
		}
		cancel()
	}()
}
