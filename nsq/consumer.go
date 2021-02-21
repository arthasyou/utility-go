package nsq

import (
	"github.com/luobin998877/go_utility/logger"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
)

// ConsumerHandler creates a service handler that will be used to handle message.
type ConsumerHandler interface {
	ProcessMessage(topic string, channel string, message []byte) error
}

var handler ConsumerHandler

// RegisterConsumerHandler with the same name, the one registered last will take effect.
func RegisterConsumerHandler(h ConsumerHandler) {
	handler = h
}

type messageHandler struct {
	topic   string
	channel string
}

func (h *messageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}
	err := handler.ProcessMessage(h.topic, h.channel, m.Body)
	if err != nil {
		logger.Error("nsq Consumer ProcessMessage error... ", zap.String("err", err.Error()))
		return err
	}
	return nil
}

var config = nsq.NewConfig()

// ConsumerStart nsq
func ConsumerStart(addr string, topic string, channel string) {
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		logger.Error("nsq Consumer Start error... ", zap.String("err", err.Error()))
	}
	consumer.AddHandler(&messageHandler{topic: topic, channel: channel})
	put(topic, consumer)

	err = consumer.ConnectToNSQLookupd(addr)
	if err != nil {
		logger.Error("nsq Consumer Lookup error... ", zap.String("err", err.Error()))
	}
}

// ConsumerStop nsq
func ConsumerStop() {
	for k, c := range m {
		c.Stop()
		delete(m, k)
	}
}

var m = make(map[string]*nsq.Consumer)

func put(topic string, consumer *nsq.Consumer) {
	m[topic] = consumer
}
