package test

import (
	"context"
	"github.com/liuxiaobleach/kafka-demo/pb"
	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"
	"log"
	"testing"
)

func TestKafka(t *testing.T) {
	w := &kafka.Writer{
		Addr: kafka.TCP("localhost:9092"),
	}

	data := pb.Data{
		Code: 1,
		Msg:  "hello world",
	}

	dataBytes, err := proto.Marshal(&data)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	msg1 := kafka.Message{
		Topic: "topic-A",
		Key:   dataBytes,
	}

	err = w.WriteMessages(context.Background(), msg1)
	if err != nil {
		log.Fatal("write error: ", err)
	}
}
