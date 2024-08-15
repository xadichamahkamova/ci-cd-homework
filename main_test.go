package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
)

func TestHealthHandler(t *testing.T) {

    router := gin.Default()
    router.GET("/health", HealthHandler)

    req, err := http.NewRequest(http.MethodGet, "/health", nil)
    if err != nil {
        t.Fatal(err)
    }

    w := httptest.NewRecorder() 
    router.ServeHTTP(w, req) 

    if status := w.Code; status != http.StatusOK {
        t.Errorf("Expected status 200 OK, got %v", status)
    }

    expected := "OK"
    if w.Body.String() != expected {
        t.Errorf("Expected response body %v, got %v", expected, w.Body.String())
    }
}
