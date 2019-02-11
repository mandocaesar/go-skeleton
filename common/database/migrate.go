package database

import (
	"github.com/jinzhu/gorm"
	authenticationModel "github.com/mandocaesar/go-skeleton/domain/authentication/model"
)

//Migrate : migrate model
func Migrate(_db *gorm.DB) {
	_db.AutoMigrate(
		&authenticationModel.User{},
	)
}

//Seed : seed database
func Seed(_db *gorm.DB) {

	tx := _db.Begin()
	var user = &authenticationModel.User{
		FirstName: "Saya User",
		LastName:  "Saya Last Name",
		Password:  "yaa password la",
		UserName:  "Username",
	}
	tx.Create(user)
	tx.Commit()
}
