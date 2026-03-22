# Go Items API

A small REST API written in Go using Gin.

This project was built as a focused onboarding exercise to demonstrate the ability to quickly become productive with Go and deliver a clean, testable, and dockerized API.

## Features

- `GET /health` returns API health status
- `GET /items` returns all items from in-memory storage
- `POST /items` creates a new item
- In-memory storage
- Unit tests
- Docker support

## Project structure

- `main.go` - application entry point
- `router.go` - route registration
- `handlers.go` - HTTP handlers
- `store.go` - in-memory storage logic
- `models.go` - request and response models
- `main_test.go` - API tests

## Requirements

- Go installed
- Docker installed (optional, only for containerized run)

## Run locally

Install dependencies and start the API:

```bash
go mod tidy
go run .