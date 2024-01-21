package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) DeleteGet(c *gin.Context) {
	c.HTML(http.StatusOK, "delete.html", gin.H{})
}
