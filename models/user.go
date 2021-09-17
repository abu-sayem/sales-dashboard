package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `gorm:"unique" json:"email"`
	Password []byte `json:"-"`
}

func (u *User) SetPassword(password string) {
	hashpassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	u.Password = hashpassword
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}