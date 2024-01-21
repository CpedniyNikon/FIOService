package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subscriber/internal/models"
	"subscriber/internal/utils/methods"
)

func (h *Handler) NewPost(c *gin.Context) {
	firstName := c.PostForm("FirstName") // получение значения параметра name из GET запроса
	lastName := c.PostForm("LastName")   // получение значения параметра email из GET запроса

	db, err := methods.InitDbConnection()
	if err != nil {
		panic(err)
	}

	p := &models.PersonData{FirstName: firstName,
		LastName: lastName}

	methods.GiveAge(p)
	methods.GiveGender(p)
	methods.GiveNationality(p)

	db.Create(p)

	c.Redirect(http.StatusFound, "/persons/")
}
