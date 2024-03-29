package main

import (
	"encoding/json"
	"fmt"
	"log"
	"publisher/internal/models"
	"publisher/internal/utils/methods"
	"time"
)

func main() {

	//res := methods.MakeRequest(fmt.Sprintf("https://api.agify.io/?name=%s", "Ruslan"))
	//resBytes := []byte(res)
	//var jsonObj map[string]interface{}
	//
	//_ = json.Unmarshal(resBytes, &jsonObj)
	//fmt.Println(jsonObj)
	//
	//age, ok := jsonObj["age"].(float64)
	//
	//if !ok {
	//	return
	//}
	//
	//fmt.Println(age)

	time.Sleep(10 * time.Second)
	nc, err := methods.InitNatsBridge()
	if err != nil {
		log.Fatalf("error initializing nats bridge: %s\n", err.Error())
	}

	res := methods.MakeRequest("https://randomuser.me/api/") // Making the request
	resBytes := []byte(res)
	var jsonObj map[string]interface{}

	_ = json.Unmarshal(resBytes, &jsonObj)

	results := jsonObj["results"].([]interface{})

	first := results[0].(map[string]interface{})

	name := first["name"].(map[string]interface{})

	p := models.Person{FirstName: name["first"].(string),
		LastName: name["last"].(string),
	}

	fmt.Println(p)
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	methods.RequestModelOrder(nc, b)

	time.Sleep(time.Millisecond)

	nc.Close()
}
