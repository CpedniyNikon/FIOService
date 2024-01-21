package methods

import (
	"encoding/json"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"subscriber/internal/models"
)

func SubscribeModelOrder(db *gorm.DB, nc stan.Conn) stan.Subscription {
	subscriber, err := nc.Subscribe("Model-person", func(m *stan.Msg) {
		var restoredOrder models.PersonData
		err := json.Unmarshal(m.Data, &restoredOrder)
		if err != nil {
			logrus.Debug(err)
		}

		restoredOrder.SetAge()
		restoredOrder.SetGender()
		restoredOrder.SetNationality()

		db.Create(&restoredOrder)
	})
	if err != nil {
		logrus.Debug(err)
	}

	return subscriber
}
