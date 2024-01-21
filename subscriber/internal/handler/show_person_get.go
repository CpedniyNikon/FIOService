package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) ShowPersonGet(c *gin.Context) {
	logrus.Debug("called show person get endpoint")
	c.HTML(http.StatusOK, "show_person.html", gin.H{})
}
