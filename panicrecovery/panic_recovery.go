package panicrecovery

import (
	"context"
	"fmt"
	"runtime/debug"

	"github.com/go-logr/logr"
)

var PanicChan = make(chan error)

// handleBubbledPanic receives errors from panicChan, then it will print them and stop() context.
func HandleBubbledPanic(ctx context.Context, stop context.CancelFunc, log logr.Logger) {
	go func() {
		for {
			select {
			case err := <-PanicChan:
				log.V(0).Error(err, "")
				stop()
			case <-ctx.Done():
				return
			}
		}
	}()
}

// panicRecovery recovers from panics and sends them to panicChan
func PanicRecovery() {
	if recovery := recover(); recovery != nil {
		PanicChan <- fmt.Errorf("[panic recovery] %s - [stack trace] %s", recovery, debug.Stack())
	}
}
