package entity

import (
	"time"
	"gorm.io/gorm"
)

// type User struct {
// 	gorm.Model
// 	FirstName string
// 	LastName  string
// 	Email     string
// 	Age       uint8
// 	BirthDay  time.Time
// }

type Admin struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	Employee []Employee `gorm:"foreignKey:AdminID"`
}

type Title struct {
	gorm.Model
	Name      string
	Employee []Employee `gorm:"foreignKey:TitleID"`
}

type Role struct {
	gorm.Model
	Name string
	Employee    []Employee `gorm:"foreignKey:RoleID"`
}

type Gender struct {
	gorm.Model
	Name      string
	Employee []Employee `gorm:"foreignKey:GenderID"`
}

type Employee struct {
	gorm.Model
	//AdminID ทำหน้าที่ FK
	AdminID *uint
	Admin   Admin
	IDCard  string `gorm:"uniqueIndex"`
	//TitleID ทำหน้าที่ FK
	TitleID *uint
	Title   Title
	Name    string
	//RoleID ทำหน้าที่ FK
	RoleID  *uint
	Role    Role
	Phonenumber string `gorm:"uniqueIndex"`
	Email       string `gorm:"uniqueIndex"`
	Password    string `json:"-"`
	//GenderID ทำหน้าที่ FK
	GenderID *uint
	Gender   Gender
	//BloodTypeID ทำหน้าที่ FK
	Salary 	uint8
	BirthDay    time.Time
}
