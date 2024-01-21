package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) NewGet(c *gin.Context) {
	logrus.Debug("called new get endpoint")
	c.HTML(http.StatusOK, "new.html", gin.H{})
}
