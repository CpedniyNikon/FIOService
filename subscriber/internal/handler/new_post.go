package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subscriber/internal/models"
	"subscriber/internal/utils/methods"
)

func (h *Handler) NewPost(c *gin.Context) {
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
