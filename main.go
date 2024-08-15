package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}

func main() {

	r := gin.Default() 

    r.GET("/health", HealthHandler) 

    r.Run(":8080") 
}
