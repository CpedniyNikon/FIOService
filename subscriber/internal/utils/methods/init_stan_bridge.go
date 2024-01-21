package methods

import (
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

func InitNatsBridge() (stan.Conn, error) {
	nc, err := stan.Connect("cluster", "subscriber", stan.NatsURL("nats://stan-server:4222"))
	if err != nil {
		logrus.Fatal(err)
	}
	return nc, err
}
