package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) NewGet(c *gin.Context) {
	c.HTML(http.StatusOK, "new.html", gin.H{})
}
