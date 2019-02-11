package model

//User : user struct
type User struct {
	UserName  string `gorm:"type:varchar(100)" valid:"alphanum, required"`
	FirstName string `gorm:"type:varchar(100)" valid:"alphanum, required"`
	LastName  string `gorm:"type:varchar(100)" valid:"alphanum, required"`
	Password  string `gorm:"type:varchar(100)" valid:"alphanum, required"`
}
