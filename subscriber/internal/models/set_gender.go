package models

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

func (Person *PersonData) SetGender() {
	res := MakeRequest(fmt.Sprintf("https://api.genderize.io/?name=%s", Person.FirstName))
	resBytes := []byte(res)
	var jsonObj map[string]interface{}

	_ = json.Unmarshal(resBytes, &jsonObj)

	gender, ok := jsonObj["gender"].(string)

	if !ok {
		logrus.Fatal("error while panic conversion")
		return
	}

	Person.Gender = gender
}
