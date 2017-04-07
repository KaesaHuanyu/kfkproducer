package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	conf := sarama.NewConfig()
	conf.Version = sarama.V0_9_0_1
	conf.Producer.Return.Successes = true
	conf.Producer.Return.Errors = true
	addresses := []string{"broker01:9092", "broker02:9092", "broker03:9092"}
	//addresses := []string{"192.168.1.129:30003", "192.168.1.129:30004", "192.168.1.129:30005"}
	producer, err := sarama.NewSyncProducer(addresses, conf)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	msg := &sarama.ProducerMessage{Topic: "daocloud", Value: sarama.StringEncoder("testing 123")}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Printf("FAILED to send message: %s\n", err)
	} else {
		log.Printf("> message sent to partition %d at offset %d\n", partition, offset)
	}
}
