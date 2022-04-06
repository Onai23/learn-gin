package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
	"learn-gin/config"
	"learn-gin/models"
)

//"github.com/jinzhu/gorm"
//	_ "github.com/jinzhu/gorm/dialects/mysql"

//set fungsi getHome() untuk route "/"
//gin.Context = membawa detail permintaan, memvalidasi dan membuat serialisasi data json
func GetHome(c *gin.Context) {
	//mengambil data dengan mengakses model
	items := []models.Article{}
	//ambil data dari table articles, kemudian simpan ke dalam variable items
	config.DB.Find(&items) /// SELECT * FROM articles
	//gin.H = set response
	c.JSON(200, gin.H{
		"status": "Berhasil ke halaman home",
		"data":   items,
	})
}

//set fungsi getArticle()
func GetArticle(c *gin.Context) {
	//get parameter title
	slug := c.Param("slug")
	var item models.Article
	//ambil data berdasarkan slug, kemudian simpan ke dalam cetakan item struct
	if config.DB.First(&item, "slug = ?", slug).RecordNotFound() {
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
func PostArticle(c *gin.Context) {
	//mengambil data yang dikirim dari form post, kirim dari url-form-encoded
	//deklarasi cetakan article untuk menambahkan databaru ke database
	item := models.Article{
		Title: c.PostForm("title"),
		Desc:  c.PostForm("desc"),
		Slug:  slug.Make(c.PostForm("title")),
	}
	//jika slug sama generate random slug
	//cek database apakah ada slug yang sama
	//jika ada, beri string random pada slug

	config.DB.Create(&item)
	//set response json
	c.JSON(200, gin.H{
		"status": "Berhasil ngepost",
		"data":   item,
	})
}
