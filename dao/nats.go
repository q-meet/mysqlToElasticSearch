package dao

import "github.com/nats-io/nats.go"
import "log"

var NatsMQ MessageQueue

func InitNats() {
	nc, err := nats.Connect(nats.DefaultURL, nats.Name("xxx"))
	if err != nil {
		log.Fatal(err)
	}
	NatsMQ = NMQImpl{NC: nc}
}

type NMQImpl struct {
	NC *nats.Conn
}

func (nmq NMQImpl) Pub(key string, data []byte) error {
	return nmq.NC.Publish(key, data)
}

func (nmq NMQImpl) Sub(key string, mh MsgHandler) error {
	_, err := nmq.NC.Subscribe(key, func(msg *nats.Msg) {
		err := mh(msg.Data)
		if err != nil {
			panic(err)
		}
	})
	return err
}
