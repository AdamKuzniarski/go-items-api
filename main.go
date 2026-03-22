package main

import "sync"

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateItemInput struct {
	Name string `json:"name" binding:"required"`
}

type Store struct {
	mu     sync.Mutex
	items  []Item
	nextID int
}
