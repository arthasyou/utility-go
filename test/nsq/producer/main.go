package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arthasyou/utility-go/nsq"
)

func main() {
	nsq.ProducerStart("127.0.0.1:4150")
	// waitExit()
	nsq.Publish("test", []byte{1, 2, 3, 4, 5, 6})
}

func waitExit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	fmt.Println("Got a signal", sig)
	now := time.Now()

	fmt.Println("Server exited", now)
}
