package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"log"
)

func main() {
	// 创建 Kafka 消费者
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		log.Fatalln("Failed to create consumer:", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("my-topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("Failed to start partition consumer:", err)
	}
	defer partitionConsumer.Close()

	// 循环接收消息
	for msg := range partitionConsumer.Messages() {
		fmt.Printf("Received message: %s\n", string(msg.Value))
	}
}
