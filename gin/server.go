package main

import "github.com/gin-gonic/gin"

func main() {
	/*
		Default returns an Engine instance with the Logger and Recovery middleware
		already attached.

		Engine is the framework's instance, it contains the muxer, middleware and
		configuration settings. Create an instance of Engine, by using New() or Default()
	*/
	server := gin.Default()

	/*
		Context is the most important part of gin. It allows us to pass variables between
		middleware, manage the flow, validate the JSON of a request and render a JSON
		response for example.
	*/
	server.GET("/test", func(ctx *gin.Context) {

		/*
			JSON serializes the given struct as JSON into the response body.
			It also sets the Content-Type as "application/json".

			H is a shortcut for map[string]interface{}
		*/
		ctx.JSON(200, gin.H{
			"message": "OK!!",
		})
	})

	server.Run(":8080")
}
