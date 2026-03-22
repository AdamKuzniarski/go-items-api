package main

import "github.com/gin-gonic/gin"

func NewRouter(handler *Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", handler.Health)
	r.GET("/items", handler.GetItems)
	r.POST("/items", handler.CreateItem)

	return r
}
