package methods

import (
	"encoding/json"
	"fmt"
	"subscriber/internal/models"
)

func GiveAge(Person *models.PersonData) {
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
