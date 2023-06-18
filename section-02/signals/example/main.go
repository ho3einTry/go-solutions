package main

import (
	"fmt"
	signals2 "github.com/ho3eintry/go-solutions/section-02/signals"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	signals := make(chan os.Signal)
	done := make(chan bool)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go signals2.CatchSig(signals, done)
	fmt.Println("Press ctrl-c to terminate...")
	<-done
	fmt.Println("Done!")
}
