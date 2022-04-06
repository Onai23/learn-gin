package main

import "github.com/gin-gonic/gin"

func main() {
	//set default route
	router := gin.Default()

	//set handler
	//gin.Context = membawa detail permintaan, memvalidasi dan membuat serialisasi data json
	router.GET("/", getHome)
	//menjalakan server
	router.Run("localhost:8080")
}

//set fungsi getHome() untuk route "/"
//gin.Context = membawa detail permintaan, memvalidasi dan membuat serialisasi data json
func getHome(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "Berhasil",
		"message": "Berhasil akses home",
	})
}
