package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) ShowPersonPost(c *gin.Context) {
	logrus.Debug("called show person post endpoint")
	PersonId := c.PostForm("PersonId")
	c.Redirect(http.StatusFound, fmt.Sprintf("/persons/%s/status", PersonId))
}
