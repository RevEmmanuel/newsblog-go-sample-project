package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	*UserDTO
}
type UserDTO struct {
	FirstName   string       `json:"firstname" validate:"required,max=30" `
	LastName    string       `json:"lastname" validate:"required,max=30"`
	Email       string       `json:"email" validate:"required,email" gorm:"unique"`
	Password    string       `json:"password" validate:"required,min=5"`
	PhoneNumber int          `json:"phonenumber" validate:"required"`
	Preferences []Preference `json:"preferences" validate:"required,validPreferences" gorm:"many2many:user_preferences"`
}
type Preference struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:user_preferences"`
}
