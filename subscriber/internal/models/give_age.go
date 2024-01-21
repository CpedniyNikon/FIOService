package models

import (
	"encoding/json"
	"fmt"
)

func (Person *PersonData) SetAge() {
	res := MakeRequest(fmt.Sprintf("https://api.agify.io/?name=%s", Person.FirstName))
	resBytes := []byte(res)
	var jsonObj map[string]interface{}

	_ = json.Unmarshal(resBytes, &jsonObj)

	age, ok := jsonObj["age"].(float64)

	if !ok {
		return
	}
	Person.Age = age
}
