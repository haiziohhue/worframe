package initialize

import (
	"github.com/Shopify/sarama"
	"log"
	"worframe/share/core"
)

type Producer struct {
	Producer   sarama.SyncProducer
	Topic      string //主题
	ProducerID int    //生产者Id
	MessageId  int
}

func SendMsg() error {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9093"}, nil)
	if err != nil {
		core.Logger.Error(err)
		return err
	}
	defer func() {
		if err = producer.Close(); err != nil {
			core.Logger.Error(err)
		}
	}()
	msg := &sarama.ProducerMessage{
		Topic: "my-topic",
		Value: sarama.StringEncoder("hello world"),
	}
	partition, offset, _ := producer.SendMessage(msg)
	log.Default().Printf("partition:%d offset:%d\n", partition, offset)
	if err != nil {
		core.Logger.Error(err)
		return err
	}
	return nil
}
