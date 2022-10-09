package life

import (
	"context"
	"os"
	"os/signal"
	"syscall"
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
