package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"subscriber/internal/models"
	"subscriber/internal/utils/methods"
)

func (h *Handler) DeletePost(c *gin.Context) {
	PersonId := c.PostForm("PersonId")

	db, err := methods.InitDbConnection()
	if err != nil {
		panic(err)
	}

	db.Delete(&models.PersonData{}, PersonId)

	c.Redirect(http.StatusFound, "/persons/")
}
