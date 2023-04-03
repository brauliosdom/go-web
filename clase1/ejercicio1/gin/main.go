package main

import (
	"github.com/gin-gonic/gin"
)

type Info struct {
	Nombre string `json: nombre`
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, map[string]any{
			"message": "pong",
		})
	})

	router.POST("/saludo", func(ctx *gin.Context) {
		var info Info
		if err := ctx.BindJSON(&info); err != nil {
			// imprime error
		}
		saludo := "Hola " + info.Nombre
		ctx.JSON(200, map[string]any{
			"saludo": saludo,
		})
	})

	router.Run(":8080")
}
