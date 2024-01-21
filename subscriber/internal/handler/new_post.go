package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"subscriber/internal/models"
	"subscriber/internal/utils/methods"
)

func (h *Handler) NewPost(c *gin.Context) {
	logrus.Debug("called new post endpoint")
	firstName := c.PostForm("FirstName")
	lastName := c.PostForm("LastName")

	db, err := methods.InitDbConnection()
	if err != nil {
		panic(err)
	}

	p := &models.PersonData{FirstName: firstName,
		LastName: lastName}

	p.SetAge()
	p.SetGender()
	p.SetNationality()

	db.Create(p)

	c.Redirect(http.StatusFound, "/persons/")
}
