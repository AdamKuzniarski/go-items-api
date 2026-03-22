package main

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateItemInput struct {
	Name string `json:"name" binding:"required"`
}
