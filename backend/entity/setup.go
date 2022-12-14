package entity

import (
  "golang.org/x/crypto/bcrypt"
  "gorm.io/gorm"
  "gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {
  return db
}

func SetupDatabase() {
  database, err := gorm.Open(sqlite.Open("se-65.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema

  database.AutoMigrate(
    &Admin{},
    &Gender{},
    &Role{},
    &Title{})

  db = database
  //ผู้ดูแลระบบ
  PasswordAdmin1, err := bcrypt.GenerateFromPassword([]byte("123456789"), 14)
	Admin1 := Admin{
		FirstName:  "Pimchanok",
		LastName:   "Somsri",
		Email:      "pimchanok447@gmail.com",
		Password:    string(PasswordAdmin1),
	}
  db.Model(&Admin{}).Create(&Admin1)

  //ตำแหน่ง
  Role1 := Role{
		Name: "Medical Physician",
	}
	db.Model(&Role{}).Create(&Role1)
  Role2 := Role{
		Name: "Nurse",
	}
	db.Model(&Role{}).Create(&Role2)
  Role3 := Role{
		Name: "Pharmacist",
	}
	db.Model(&Role{}).Create(&Role3)
  Role4 := Role{
		Name: "Finance Officer",
	}
	db.Model(&Role{}).Create(&Role4)
  Role5 := Role{
		Name: "Medical Record Officer",
	}
	db.Model(&Role{}).Create(&Role5)
  
  //คำนำหน้า
  Title1 := Title{
		Name: "Mr.",
	}
	db.Model(&Title{}).Create(&Title1)
  Title2 := Title{
		Name: "Mrs.",
	}
	db.Model(&Title{}).Create(&Title2)
  Title3 := Title{
		Name: "Ms.",
	}
	db.Model(&Title{}).Create(&Title3)
  Title4 := Title{
		Name: "DR.",
	}
	db.Model(&Title{}).Create(&Title4)
  Title5 := Title{
		Name: "Pharm.",
	}
	db.Model(&Title{}).Create(&Title5)

}