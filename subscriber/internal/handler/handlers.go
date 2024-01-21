package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	routes *gin.Engine
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitAuthRoutes() *gin.Engine {
	h.routes = gin.New()
	h.routes.Use(cors.Default())
	h.routes.LoadHTMLGlob("templates/*")
	orders := h.routes.Group("/persons")
	{
		orders.GET("/", h.Get)
		orders.GET("/new", h.NewGet)
		orders.POST("/new", h.NewPost)
		orders.GET("/:personId/status", h.GetPersonById)
	}

	return h.routes
}
