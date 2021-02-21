package nsq

import (
	"github.com/luobin998877/go_utility/logger"
	"github.com/nsqio/go-nsq"
	"go.uber.org/zap"
)

var p *nsq.Producer

// ProducerStart nsq
func ProducerStart(addr string) {
	producer, err := nsq.NewProducer(addr, config)
	if err != nil {
		logger.Error("nsq Producer Starting error... ", zap.String("err", err.Error()))
	}
	p = producer

}

// ProducerStop nsq
func ProducerStop() {
	p.Stop()
}

// Publish to nsq Producer
func Publish(topic string, message []byte) {
	p.Publish(topic, message)
}
