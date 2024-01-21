package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) DeleteGet(c *gin.Context) {
	logrus.Debug("called delete get endpoint")
	c.HTML(http.StatusOK, "delete.html", gin.H{})
}
