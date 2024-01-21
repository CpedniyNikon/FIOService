package models

import (
	"fmt"
	"io"
	"net/http"
)

func MakeRequest(URL string) string {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", URL, nil)
	req.Header.Set("Header_Key", "Header_Value")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Err is", err)
	}
	defer res.Body.Close()

	resBody, _ := io.ReadAll(res.Body)
	response := string(resBody)

	return response
}
