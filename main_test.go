package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return setupRouter(NewStore())
}

func TestHealth(t *testing.T) {
	router := setupTestRouter()

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}
}

func TestPostItem(t *testing.T) {
	router := setupTestRouter()

	body := map[string]string{
		"name": "Keyboard",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/items", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w.Code)
	}

	var item Item
	if err := json.Unmarshal(w.Body.Bytes(), &item); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if item.Name != "Keyboard" {
		t.Fatalf("expected item name Keyboard, got %s", item.Name)
	}
}

func TestGetItemsAfterPost(t *testing.T) {
	router := setupTestRouter()

	body := map[string]string{
		"name": "Mouse",
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal body: %v", err)
	}

	postReq := httptest.NewRequest(http.MethodPost, "/items", bytes.NewBuffer(jsonBody))
	postReq.Header.Set("Content-Type", "application/json")
	postW := httptest.NewRecorder()
	router.ServeHTTP(postW, postReq)

	if postW.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", postW.Code)
	}

	getReq := httptest.NewRequest(http.MethodGet, "/items", nil)
	getW := httptest.NewRecorder()
	router.ServeHTTP(getW, getReq)

	if getW.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", getW.Code)
	}

	var items []Item
	if err := json.Unmarshal(getW.Body.Bytes(), &items); err != nil {
		t.Fatalf("failed to unmarshal response: %v", err)
	}

	if len(items) != 1 {
		t.Fatalf("expected 1 item, got %d", len(items))
	}

	if items[0].Name != "Mouse" {
		t.Fatalf("expected item name Mouse, got %s", items[0].Name)
	}
}
