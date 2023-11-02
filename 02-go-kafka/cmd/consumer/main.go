package main

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": "gokafka_kafka_1:9092",
		"client.id":         "goapp-consumer", // aplicação que consome o tópico
		"group.id":          "goapp-group2",   // grupo de consumo
		"auto.offset.reset": "earliest",       // pega as mensagens des do ínicio
	}
	c, err := kafka.NewConsumer(configMap) // consumer
	if err != nil {
		fmt.Println("error consumer", err.Error())
	}
	topics := []string{"teste"} // tópico que vai ser consumido
	c.SubscribeTopics(topics, nil)
	for {
		msg, err := c.ReadMessage(-1) // fica conectado para sempre
		if err == nil {
			fmt.Println(string(msg.Value), msg.TopicPartition) // se der erro imprime o erro, mensagem e partição
		}
	}
}
