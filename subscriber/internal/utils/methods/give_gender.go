package methods

import (
	"encoding/json"
	"fmt"
	"subscriber/internal/models"
)

func GiveGender(Person *models.PersonData) {
	res := MakeRequest(fmt.Sprintf("https://api.genderize.io/?name=%s", Person.FirstName))
	resBytes := []byte(res)
	var jsonObj map[string]interface{}

	_ = json.Unmarshal(resBytes, &jsonObj)

	gender, ok := jsonObj["gender"].(string)

	if !ok {
		return
	}

	Person.Gender = gender
}
