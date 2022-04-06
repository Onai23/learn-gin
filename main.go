package main

import "github.com/gin-gonic/gin"

func main() {
	//set default route
	r := gin.Default()

	//set handler
	//gin.Context = membawa detail permintaan, memvalidasi dan membuat serialisasi data json
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "berhasil akses home",
		})
	})
	//menjalakan server
	r.Run("localhost:8080")
}
