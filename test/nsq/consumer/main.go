package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arthasyou/utility-go/nsq"
)

type consumerHandler struct{}

func (h *consumerHandler) ProcessMessage(topic string, channel string, message []byte) error {
	fmt.Println("topic: ", topic, "channel: ", channel, "message: ", message)
	return nil
}

func main() {
	nsq.RegisterConsumerHandler(&consumerHandler{})
	nsq.ConsumerStart("localhost:4161", "test", "channel")
	waitExit()
}

func waitExit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	fmt.Println("Got a signal", sig)
	now := time.Now()

	fmt.Println("Server exited", now)
}
