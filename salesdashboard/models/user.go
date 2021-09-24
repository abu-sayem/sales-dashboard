package models

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        uint          `json:"id"`
	FirstName string 				`json:"first_name"`
	LastName  string  			`json:"last_name"`
	Email     string     		`gorm:"unique" json:"email"`
	Password  []byte 				`json:"-"`
	RoleId    uint     			`json:"role_id"`
	Role      Role          `json:"role" gorm:"foreignKey:RoleId"`
}

func (u *User) SetPassword(password string) {
	hashpassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	u.Password = hashpassword
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}

func (u *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Product{}).Count(&total)
	return total
}

func (u *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var users []User

	db.Preload("Role").Offset(offset).Limit(limit).Find(&users)
	
	return users
}