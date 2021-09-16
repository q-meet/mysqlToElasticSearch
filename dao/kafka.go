package dao

import (
	"context"
	"fmt"
	kafka "github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"log"
)

var KafkaMQ KafkaMessageQueue

func InitKafka() {
	topic := fmt.Sprintf("%v", viper.Get("kafka.topic"))
	partition := 0
	address := fmt.Sprintf("%v:%v", viper.Get("kafka.host"), viper.Get("kafka.port"))

	conn, err := kafka.DialLeader(context.Background(), "tcp", address, topic, partition)
	if err != nil {
		panic(err.Error())
	}
	/*
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
	}


		r := kafka.NewReader(kafka.ReaderConfig{
			Brokers:   []string{address},
			Topic:     topic,
			Dialer:    dialer,
			MinBytes:  10e3, // 10KB
			MaxBytes:  10e6, // 10MB
			Partition: partition,
		})

		w := kafka.NewWriter(kafka.WriterConfig{
			Brokers:  []string{address},
			Topic:    topic,
			Balancer: &kafka.Hash{},
			Dialer:   dialer,
		})
	*/
	KafkaMQ = KafkaMQImpl{
		//R: r,
		//W: w,
		Conn: conn,
	}
}

//R *kafka.Reader
//W *kafka.Writer
type KafkaMQImpl struct {
	Conn *kafka.Conn
}

func (nmq KafkaMQImpl) Clear() error {
	//关闭连接时机
	return nmq.Conn.Close()
}

func (nmq KafkaMQImpl) Pub(key string, data []byte) error {
	_, err := nmq.Conn.WriteMessages(
		kafka.Message{
			Key:   []byte(key),
			Value: data,
		},
	)

	if err != nil {
		log.Fatal("failed to write messages:", err)
	}
	return nil
}

func (nmq KafkaMQImpl) Sub(key string, mh MsgHandler) error {

	for {
		m, err := nmq.Conn.ReadMessage(10e6)

		if err != nil {
			break
		}

		// TODO: process message
		fmt.Printf("message at offset %d: %s = %s %v\n", m.Offset, string(m.Key), string(m.Value), string(m.Key) == key)

		if string(m.Key) == key {
			err = mh(m.Value)
			if err != nil {
				panic(err)
			}
		}
	}

	if err := nmq.Conn.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
	return nil
}
