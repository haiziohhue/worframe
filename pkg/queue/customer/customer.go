package customer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"worframe/share/core"
)

func Consumer() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9093"}, sarama.NewConfig())
	if err != nil {
		core.Logger.Error(err)

		return
	}
	defer func() {
		if err = consumer.Close(); err != nil {
			core.Logger.Error(err)
			return
		}
	}()
	// 获取消费者的分片接口，sarama.OffsetNewest 标识获取新的消息
	partitionConsumer, err := consumer.ConsumePartition("topic", 0, sarama.OffsetNewest)
	if err != nil {
		core.Logger.Error(err)
		return
	}
	defer func() {
		if err = partitionConsumer.Close(); err != nil {
			core.Logger.Error(err)
			return
		}
	}()
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("分片:%d 偏移:%d Key:%s Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
	}
}
