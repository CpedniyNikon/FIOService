package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"subscriber/internal/models"
	"subscriber/internal/utils/methods"
)

func (h *Handler) UpdatePost(c *gin.Context) {
	logrus.Debug("called update post endpoint")
	PersonId := c.PostForm("PersonId")
	firstName := c.PostForm("FirstName")
	lastName := c.PostForm("LastName")

	db, err := methods.InitDbConnection()
	if err != nil {
		logrus.Panic("called update get endpoint")
		panic(err)
	}

	db.Model(&models.PersonData{}).Where("Id = ?", PersonId).Update("FirstName", firstName)
	db.Model(&models.PersonData{}).Where("Id = ?", PersonId).Update("LastName", lastName)

	c.Redirect(http.StatusFound, "/persons/")
}
