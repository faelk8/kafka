package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	deliveryChan := make(chan kafka.Event) // canal para ler o evento
	producer := NewKafkaProducer()
	Publish("transferiu", "teste", producer, []byte("transferecia2"), deliveryChan) // a chave server para ir para a mesma partição
	DeliveryReport(deliveryChan)                                                    // async

	//e := <-deliveryChan // em quanto a mensagem não chegar fica travado aqui
	//msg := e.(*kafka.Message) // pega a mensagem do evento
	//if msg.TopicPartition.Error != nil { // tratamento do erro
	//	fmt.Println("Erro ao enviar")
	//} else {
	//	fmt.Println("Mensagem enviada:", msg.TopicPartition) // mostra a mensagem
	//}
	//

}

func NewKafkaProducer() *kafka.Producer {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers":   "gokafka_kafka_1:9092",
		"delivery.timeout.ms": "0",    // fica aguardando para sempre a mensagem
		"acks":                "all",  // o lider e o broker recebe as menssagens
		"enable.idempotence":  "true", // indepotente, mensagem entregue na ordem e pelomenos uma vez os acks acknologe precisa ser all
	}
	p, err := kafka.NewProducer(configMap)
	if err != nil {
		log.Println(err.Error())
	}
	return p
}

// cria uma mensagem forcando uma string - canal de kafka - toda mensagem publica no cal
func Publish(msg string, topic string, producer *kafka.Producer, key []byte, deliveryChan chan kafka.Event) error {
	message := &kafka.Message{
		Value:          []byte(msg),                                                        //conteudo da mensagem - forma de byte
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny}, //topico e particao
		Key:            key,                                                                // chave de indentificao
	}
	err := producer.Produce(message, deliveryChan) // produzindo a menssagem
	if err != nil {                                // tratamento do erro
		return err
	}
	return nil
}

func DeliveryReport(deliveryChan chan kafka.Event) { // loop infinito que fica pegando a mensagem
	for e := range deliveryChan {
		switch ev := e.(type) { // tipo do evento
		case *kafka.Message: // kafka mesage
			if ev.TopicPartition.Error != nil {
				fmt.Println("Erro ao enviar")
			} else {
				fmt.Println("Mensagem enviada:", ev.TopicPartition)
				// anotar no banco de dados que a mensagem foi processado.
				// ex: confirma que uma transferencia bancaria ocorreu.
			}
		}
	}
}
