package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"subscriber/internal/models"
	"subscriber/internal/utils/methods"
)

func (h *Handler) GetPersonById(c *gin.Context) {
	logrus.Debug("called get person endpoint")
	personId := c.Param("personId")

	db, err := methods.InitDbConnection()
	if err != nil {
		panic(err)
	}

	var order models.PersonData
	db.First(&order, personId)

	jsonData, err := json.MarshalIndent(order, "", "  ")
	if err != nil {
		logrus.Debug("Error:", err)
	}

	c.HTML(http.StatusOK, "get.html", gin.H{
		"jsonString": string(jsonData),
	})
}
