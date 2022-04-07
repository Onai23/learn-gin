package models

import "github.com/jinzhu/gorm"

//Model, dibentuk database orm, mewakili tabel2 yang ada di database
//id unique akan otomatis dibuatkan
//deklarasi cetakan User, user dapat memiliki lebih dari 1 article
type User struct {
	gorm.Model
	Articles []Article //1 user dapat memiliki lebih dari 1 article
	Username string
	FullName string
	Email    string
	SocialId string
	Provider string
	Avatar   string
	Role     bool `gorm:"default:0"`
}
