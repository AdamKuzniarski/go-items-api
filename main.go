package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter(store *Store) *gin.Engine {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/items", func(c *gin.Context) {
		store.mu.Lock()
		defer store.mu.Unlock()

		c.JSON(http.StatusOK, store.items)
	})

	r.POST("/items", func(c *gin.Context) {
		var input CreateItemInput

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		store.mu.Lock()
		defer store.mu.Unlock()

		item := Item{
			ID:   store.nextID,
			Name: input.Name,
		}

		store.items = append(store.items, item)
		store.nextID++

		c.JSON(http.StatusCreated, item)
	})

	return r
}
func main() {
	store := NewStore()
	r := setupRouter(store)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
