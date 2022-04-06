package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
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

//deklarasi variable global tipe *gorm.DB
var DB *gorm.DB

func main() {
	//deklarasi variable err tipe error
	var err error
	//set koneksi database
	DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/learngin?charset=utf8mb4&parseTime=True&loc=Local")
	//cek jika ada error
	if err != nil {
		panic("failed to connect database")
	}
	defer DB.Close()

	//memigrasi cetakan Article sebagai tabel di database
	DB.AutoMigrate(&Article{})

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
			articles.GET("/:slug", getArticle)
			articles.POST("/", postArticle)
		}
	}

	//menjalakan server
	router.Run("localhost:8080")
}

//set fungsi getHome() untuk route "/"
//gin.Context = membawa detail permintaan, memvalidasi dan membuat serialisasi data json
func getHome(c *gin.Context) {
	//mengambil data dengan mengakses model
	items := []Article{}
	//ambil data dari table articles, kemudian simpan ke dalam variable items
	DB.Find(&items) /// SELECT * FROM articles
	//gin.H = set response
	c.JSON(200, gin.H{
		"status": "Berhasil ke halaman home",
		"data":   items,
	})
}

//set fungsi getArticle()
func getArticle(c *gin.Context) {
	//get parameter title
	slug := c.Param("slug")
	var item Article
	//ambil data berdasarkan slug, kemudian simpan ke dalam cetakan item struct
	if DB.First(&item, "slug = ?", slug).RecordNotFound() {
		//set response
		c.JSON(404, gin.H{"status": "error", "message": "record not found"})
		c.Abort() //hentikan request
		return
	}
	//set response dari parameter
	c.JSON(200, gin.H{
		"status":  "Berhasil",
		"message": item,
	})
}

//set fungsi postArticle()
//gin.Context = membawa detail permintaan, memvalidasi dan membuat serialisasi data json
func postArticle(c *gin.Context) {
	//mengambil data yang dikirim dari form post, kirim dari url-form-encoded
	//deklarasi cetakan article untuk menambahkan databaru ke database
	item := Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}
	//jika slug sama generate random slug
	//cek database apakah ada slug yang sama
	//jika ada, beri string random pada slug

	DB.Create(&item)
	//set response json
	c.JSON(200, gin.H{
		"status": "Berhasil ngepost",
		"data":   item,
	})
}
