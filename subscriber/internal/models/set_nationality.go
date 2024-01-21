package models

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

func (Person *PersonData) SetNationality() {
	fmt.Println(Person)
	res := MakeRequest(fmt.Sprintf("https://api.nationalize.io/?name=%s", Person.FirstName))
	resBytes := []byte(res)
	var jsonObj map[string]interface{}

	_ = json.Unmarshal(resBytes, &jsonObj)

	countries, ok := jsonObj["country"].([]interface{})

	if !ok || len(countries) == 0 {
		logrus.Fatal("error while panic conversion")
		return
	}

	country, ok := countries[0].(map[string]interface{})["country_id"].(string)

	if !ok {
		logrus.Fatal("error while panic conversion")
		return
	}

	Person.Nationality = country
}
