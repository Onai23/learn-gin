package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Model, dibentuk database orm, mewakili tabel2 yang ada di database
//id unique akan otomatis dibuatkan
//deklarasi cetakan Article
type Article struct {
	gorm.Model
	Title string
	Slug  string `gorm:"unique_index"`
	Desc  string `sql:"type:text"`
}

func main() {
	//set koneksi database
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/learngin?charset=utf8mb4&parseTime=True&loc=Local")
	//cek jika ada error
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	//memigrasi cetakan Article sebagai tabel di database
	db.AutoMigrate(&Article{})

	//set default route
	router := gin.Default()
	//set group dari router ke /api/vi
	v1 := router.Group("/api/v1")
	{
		//membuat route khusus menangani article, -> ke home, post article, get article
		articles := v1.Group("/articles")
		{
			//set handler
			articles.GET("/", getHome)
			articles.GET("/:title", getArticle)
			articles.POST("/", postArticle)
		}
		//membuat router khusus menangani users,
		//users := v1.Group("/users")
		//{
		//	//set handler
		//	users.GET("/", getUser)
		//}
	}

	//menjalakan server
	router.Run("localhost:8080")
}

//set fungsi getHome() untuk route "/"
//gin.Context = membawa detail permintaan, memvalidasi dan membuat serialisasi data json
func getHome(c *gin.Context) {
	//gin.H = set response
	c.JSON(200, gin.H{
		"status":  "Berhasil",
		"message": "Berhasil akses home",
	})
}

//set fungsi getArticle()
func getArticle(c *gin.Context) {
	//get parameter title
	title := c.Param("title")
	//set response dari parameter
	c.JSON(200, gin.H{
		"status":  "Berhasil",
		"message": title,
	})
}

//set fungsi postArticle()
//gin.Context = membawa detail permintaan, memvalidasi dan membuat serialisasi data json
func postArticle(c *gin.Context) {
	//mengambil data yang dikirim dari form post, kirim dari url-form-encoded
	title := c.PostForm("title")
	desc := c.PostForm("desc")

	//set response json
	c.JSON(200, gin.H{
		"title":  title,
		"desc":   desc,
		"status": "Berhasil ngepost",
	})
}
