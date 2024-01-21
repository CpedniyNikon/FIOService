package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"subscriber/internal/models"
	"subscriber/internal/utils/methods"
)

func (h *Handler) GetPersonById(c *gin.Context) {
	personId := c.Param("personId")

	db, err := methods.InitDbConnection()
	if err != nil {
		panic(err)
	}

	var order models.PersonData
	db.First(&order, personId)

	jsonData, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.HTML(http.StatusOK, "get.html", gin.H{
		"jsonString": string(jsonData),
	})
}
