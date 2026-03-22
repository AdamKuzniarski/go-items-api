package main

import "sync"

type Store struct {
	mu     sync.RWMutex
	items  []Item
	nextID int
}

func NewStore() *Store {
	return &Store{
		items:  []Item{},
		nextID: 1,
	}
}

func (s *Store) ListItems() []Item {
	s.mu.RLock()
	defer s.mu.RUnlock()

	itemsCopy := make([]Item, len(s.items))
	copy(itemsCopy, s.items)

	return itemsCopy
}

func (s *Store) AddItem(name string) Item {
	s.mu.Lock()
	defer s.mu.Unlock()

	item := Item{
		ID:   s.nextID,
		Name: name,
	}

	s.items = append(s.items, item)
	s.nextID++

	return item
}
