package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Type user struct
type User struct {
	//gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;"`
	Nama     string    `json:"nama"`
	Email    string    `json:"email"`
	Alamat   string    `json:"alamat"`
	Jurusan  string    `json:"jurusan"`
	Gender   string    `json:"gender"`
	Semester int       `json:"semester"`
	IPK      float32   `json:"ipk"`
}

// Users struct : Field untuk menyimpan array dari User
type Users struct {
	Users []User `json:"users"`
}

// buat uuid baru sebelum buat record baru
func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4 sebagai ID
	user.ID = uuid.New()
	return
}
