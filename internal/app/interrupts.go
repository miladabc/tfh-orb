package app

import (
	"os"
	"os/signal"
	"syscall"
)

func handleInterrupts() <-chan os.Signal {
	signals := make(chan os.Signal, 1)
	signal.Notify(
		signals,
		syscall.SIGTERM,
		syscall.SIGINT,
	)

	return signals
}
