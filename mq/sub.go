package mq

import (
	"context"
	"os"
	"strconv"
	"time"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

type kafkaMsgHandler func(m kafka.Message) error

func RegisterSubscriber(kafkaTopic string, handleFunc kafkaMsgHandler) {

	maxWait, _ := strconv.Atoi(os.Getenv("MAX_WAIT"))

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{os.Getenv("BOOTSTRAP_SERVER")},
		GroupTopics: []string{kafkaTopic},
		GroupID:     os.Getenv("CONSUMER_GROUPID"),
		//Topic:    ,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  time.Second * time.Duration(maxWait),
	})
	r.SetOffset(kafka.LastOffset)
	KafkaConsumerList = append(KafkaConsumerList, r)

	for {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			break
		}
		if err := handleFunc(m); err != nil {
			log.Warnf("Handler func failed: %s", err.Error())
		}
		r.CommitMessages(context.Background(), m)
	}

}

// fetch message from specific topic and throw it to a kafka.Message channel
func SubscribeKafkaTopic(kafkaTopic string, ch chan kafka.Message) {

	maxWait, _ := strconv.Atoi(os.Getenv("MAX_WAIT"))

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{os.Getenv("BOOTSTRAP_SERVER")},
		GroupTopics: []string{kafkaTopic},
		GroupID:     os.Getenv("CONSUMER_GROUPID"),
		//Topic:    ,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		MaxWait:  time.Second * time.Duration(maxWait),
	})
	r.SetOffset(kafka.LastOffset)
	KafkaConsumerList = append(KafkaConsumerList, r)

	for {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			break
		}
		ch <- m
		r.CommitMessages(context.Background(), m)
	}

}

func SimpleLogHandler(m kafka.Message) error {
	log.Printf("message at offset %d: %v = %v\n", m.Offset, string(m.Key), string(m.Value))
	return nil
}

// func StrategyEventHandler(m kafka.Message) error {
// 	var strategyEvent StrategyEvent
// 	err := json.Unmarshal(m.Value, &strategyEvent)
// 	if err != nil {
// 		return err
// 	}

// 	log.Infof("Received event from topic strategy, event id: %s", strategyEvent.EventId)

// 	switch strategyEvent.Action {
// 	case "create":
// 		err = models.CreateStrategy(strategyEvent.UserId, strategyEvent.StrategyName)

// 	case "delete":
// 		_, err = models.DeleteUserStrategy(strategyEvent.UserId, strategyEvent.StrategyName)

// 	case "update":
// 		_, err = models.UpdateUserStrategy(strategyEvent.UserId, strategyEvent.StrategyName, strategyEvent.StrategyUpdateDetail)

// 	default:
// 		log.Info("Recieved an strategy event that is not create/delete/update")
// 	}

// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
