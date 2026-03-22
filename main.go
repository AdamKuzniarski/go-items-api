package main

func main() {
	store := NewStore()
	handler := NewHandler(store)
	router := NewRouter(handler)

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
