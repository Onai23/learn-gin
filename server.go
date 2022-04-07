package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"learn-gin/config"
	"learn-gin/routes"
)

func main() {
	config.InitDB()
	defer config.DB.Close()
	//memanggil method dari gotenv
	gotenv.Load()
	//set default route
	router := gin.Default()
	//set group dari router ke /api/vi
	v1 := router.Group("/api/v1")
	{
		//menambahkan path parameter /auth/:provider
		v1.GET("auth/:provider", routes.RedirectHandler)
		v1.GET("/auth/:provider/callback", routes.CallbackHandler)

		//membuat route khusus menangani article, -> ke home, post article, get article
		articles := v1.Group("/articles")
		{
			//set handler
			articles.GET("/", routes.GetHome)
			articles.GET("/:slug", routes.GetArticle)
			articles.POST("/", routes.PostArticle)
		}
	}

	//menjalakan server
	router.Run("localhost:8080")

}
