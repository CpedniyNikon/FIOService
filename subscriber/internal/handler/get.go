package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"subscriber/internal/models"
	"subscriber/internal/utils/methods"
)

func (h *Handler) Get(c *gin.Context) {
	logrus.Debug("called get endpoint")
	db, err := methods.InitDbConnection()
	if err != nil {
		panic(err)
	}

	var persons []models.PersonData
	db.Find(&persons)

	jsonData, err := json.MarshalIndent(persons, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.HTML(http.StatusOK, "get.html", gin.H{
		"jsonString": string(jsonData),
	})
}
