package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"learn-gin/models"
)

//deklarasi variable global tipe *gorm.DB
var DB *gorm.DB

func InitDB() {
	//deklarasi variable err tipe error
	var err error
	//set koneksi database
	DB, err = gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/learngin?charset=utf8mb4&parseTime=True&loc=Local")
	//cek jika ada error
	if err != nil {
		panic("failed to connect database")
	}

	//memigrasi cetakan User sebagai tabel di database
	DB.AutoMigrate(&models.User{})
	//memigrasi cetakan Article sebagai tabel di database
	DB.AutoMigrate(&models.Article{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	//migrasi relasi one to many user dengan article
	DB.Model(&models.User{}).Related(&models.Article{})

}
