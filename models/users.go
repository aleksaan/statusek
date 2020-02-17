package models

import "github.com/jinzhu/gorm"

// User - Add fields `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt`
type User struct {
	gorm.Model
	UserName  string
	UserCards []UserCards `gorm:"foreignkey:UserRefer"`
}

func (user *User) TableName() string {
	// custom table name, this is default
	return "statuses.users"
}

//UserCards -
type UserCards struct {
	gorm.Model
	CardNumber string
	UserRefer  uint
}

func (cards *UserCards) TableName() string {
	// custom table name, this is default
	return "statuses.cards"
}
