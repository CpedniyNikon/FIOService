package methods

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"gorm.io/gorm"
	"subscriber/internal/models"
)

func SubscribeModelOrder(db *gorm.DB, nc stan.Conn) stan.Subscription {
	subscriber, err := nc.Subscribe("Model-person", func(m *stan.Msg) {
		var restoredOrder models.PersonData
		err := json.Unmarshal(m.Data, &restoredOrder)
		if err != nil {
			fmt.Println(err)
		}

		restoredOrder.SetAge()
		restoredOrder.SetGender()
		restoredOrder.SetNationality()

		db.Create(&restoredOrder)
	})
	if err != nil {
		fmt.Println(err)
	}

	return subscriber
}
