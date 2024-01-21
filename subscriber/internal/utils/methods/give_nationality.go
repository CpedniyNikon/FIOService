package methods

import (
	"encoding/json"
	"fmt"
	"subscriber/internal/models"
)

func GiveNationality(Person *models.PersonData) {
	fmt.Println(Person)
	res := MakeRequest(fmt.Sprintf("https://api.nationalize.io/?name=%s", Person.FirstName))
	resBytes := []byte(res)
	var jsonObj map[string]interface{}

	_ = json.Unmarshal(resBytes, &jsonObj)

	countries, ok := jsonObj["country"].([]interface{})

	if !ok || len(countries) == 0 {
		return
	}

	country, ok := countries[0].(map[string]interface{})["country_id"].(string)

	if !ok {
		return
	}

	Person.Nationality = country
}
