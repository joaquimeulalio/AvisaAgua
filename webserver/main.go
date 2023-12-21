package main

import (
	"abobora/motorstatus"

	"github.com/gin-gonic/gin"
)

func main() {
	motorstatus.Teste()
	router := gin.Default()
	router.GET("/avisaagua/teste", teste)
	router.GET("/avisaagua/leitura", motorstatus.Leitura)
	router.Run(":8080")
}

func teste(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "teste avisaAgua!",
	})
}
