package models

import "github.com/jinzhu/gorm"

//Model, dibentuk database orm, mewakili tabel2 yang ada di database
//id unique akan otomatis dibuatkan
//deklarasi cetakan Article
type Article struct {
	gorm.Model
	Title string
	Slug  string `gorm:"unique_index"`
	Desc  string `sql:"type:text"`
}
