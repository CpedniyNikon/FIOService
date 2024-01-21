package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) UpdateGet(c *gin.Context) {
	logrus.Debug("called update get endpoint")
	c.HTML(http.StatusOK, "update.html", gin.H{})
}
