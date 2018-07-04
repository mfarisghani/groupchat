package nsq

import (
	"log"

	"github.com/groupchat/chat"
	"github.com/nsqio/go-nsq"
)

type Publisher struct {
	producer *nsq.Producer
}

func NewPublisher(addr string) (chat.Publisher, error) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(addr, config)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	publisher := &Publisher{
		producer: producer,
	}

	return publisher, nil
}

func (p *Publisher) Publish(roomID chat.RoomID, message string) error {
	if err := p.producer.Publish(string(roomID), []byte(message)); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
