package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	store *Store
}

func NewHandler(store *Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *Handler) GetItems(c *gin.Context) {
	c.JSON(http.StatusOK, h.store.ListItems())
}

func (h *Handler) CreateItem(c *gin.Context) {
	var input CreateItemInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	item := h.store.AddItem(input.Name)
	c.JSON(http.StatusCreated, item)
}
