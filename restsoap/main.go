package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/consume/rest", consumeREST) // REST API endpoint
	r.GET("/consume/soap", consumeSOAP) // SOAP API endpoint

	r.Run(":8080")
}
