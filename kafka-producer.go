package main

import (
	"github.com/Shopify/sarama"
	"log"
	"strconv"
	"time"
)

var (
	kafkaBroker = []string{"localhost:9092"}
	KafkaTopic  = "payments"
)

func ProduceEvent() {

	producer, err := setupProducer()
	if err != nil {
		panic(err)
	} else {
		log.Println("Kafka AsyncProducer up and running!")
	}

	// Will Produce Events
	produceMessages(producer)

	log.Printf("Kafka AsyncProducer finished produced.")
}

// setupProducer will create a AsyncProducer and returns it
func setupProducer() (sarama.AsyncProducer, error) {

	// We can do some stuff here like retry mechanism
	config := sarama.NewConfig()

	// NewAsyncProducer creates a new AsyncProducer using the given broker addresses and configuration.
	return sarama.NewAsyncProducer(kafkaBroker, config)
}

// produceMessages will send '0-9' to KafkaTopic each second, until goes i 10
func produceMessages(producer sarama.AsyncProducer) {

	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		message := &sarama.ProducerMessage{
			Topic: KafkaTopic,
			Value: sarama.StringEncoder(strconv.Itoa(i))}

		select {
		case producer.Input() <- message:
			log.Println("New Message produced value:", i)
		}
	}
	producer.AsyncClose() // Trigger a shutdown of the producer.
}

/*
TODO:
	//Creating a Kafka Topic
	// Listing Topics
	//Delete Kafka topic
	// Publish Message
	// Consume Message
*/
