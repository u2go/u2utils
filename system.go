package u2utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func GracefulShutdown(callback func(os.Signal)) {
	// wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	sig := <-ch
	fmt.Println("Received signal, shutting down...")
	callback(sig)
}
