package mq

import (
	"context"

	kafka "github.com/segmentio/kafka-go"
	log "github.com/sirupsen/logrus"
)

func PublishMsgNoKey(topic string, payload []byte) error {
	err := KafkaProducer.WriteMessages(context.Background(), kafka.Message{
		Topic: topic,
		Value: payload,
	})

	if err != nil {
		log.Error(err.Error())
		log.Warnf("Error in publishing to topic %s, resending to retry queue", topic)
		PublishMsgRetryQ(payload)
		return err
	}

	log.Infof("Successfully published event to topic %s", topic)
	return nil
}

func PublishMsgRetryQ(payload []byte) error {
	err := KafkaProducer.WriteMessages(context.Background(), kafka.Message{
		Topic: "retry_queue",
		Value: payload,
	})
	if err != nil {
		//log.Error
		return err
	}
	return nil
}
