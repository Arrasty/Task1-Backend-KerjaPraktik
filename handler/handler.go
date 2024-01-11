package handler

import (
	"github.com/Arrasty/tugas/database"
	"github.com/Arrasty/tugas/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateUser: Fungsi untuk membuat pengguna baru dalam database
func CreateUser(c *fiber.Ctx) error {
	db := database.DB.Db    //akses objek dari instance DB yg telah diinisialisasikan sebelumnya
	user := new(model.User) //buat instance baru

	// Membaca body request dan menyimpannya dalam objek user
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	// Membuat entitas user baru dalam database
	err = db.Create(&user).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create user", "data": err})
	}
	// Return respon the created user
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "User has created", "data": user})
}

// Get All Users from db untuk mendapatkan semua pengguna dari database
func GetAllUsers(c *fiber.Ctx) error {
	db := database.DB.Db   // akses database dari instance DB
	var users []model.User //var user slice untuk nampung objek dari model user

	// find all users in the database
	db.Find(&users)
	// If no user found, return an error
	if len(users) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Users not found", "data": nil})
	}
	// return users if there is no error
	return c.Status(200).JSON(fiber.Map{"status": "sucess", "message": "Users Found", "data": users})
}

// GetSingleUser dari database
func GetSingleUser(c *fiber.Ctx) error {
	db := database.DB.Db
	// get id dari parameter
	id := c.Params("id")
	var user model.User // user sebagai objek model user yang menampung data

	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error bang", "message": "User not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User Found", "data": user})
}

// update a user in db
func UpdateUser(c *fiber.Ctx) error {
	type updateUser struct {
		Nama     string  `json:"nama"`
		Email    string  `json:"email"`
		Alamat   string  `json:"alamat"`
		Jurusan  string  `json:"jurusan"`
		Gender   string  `json:"gender"`
		Semester int     `json:"semester"`
		IPK      float32 `json:"ipk"`
	}
	db := database.DB.Db
	var user model.User
	// get id dari parameter
	id := c.Params("id")

	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	//Mendeklarasikan variabel updateUserData untuk menyimpan data pembaruan dari request body
	var updateUserData updateUser
	// Mengurai data request body ke dalam struct updateUserData dan cek error
	err := c.BodyParser(&updateUserData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}

	// Update tiap fields
	if updateUserData.Nama != "" {
		user.Nama = updateUserData.Nama
	}
	if updateUserData.Email != "" {
		user.Email = updateUserData.Email
	}
	if updateUserData.Alamat != "" {
		user.Alamat = updateUserData.Alamat
	}
	if updateUserData.Jurusan != "" {
		user.Jurusan = updateUserData.Jurusan
	}
	if updateUserData.Gender != "" {
		user.Gender = updateUserData.Gender
	}
	if updateUserData.Semester != 0 {
		user.Semester = updateUserData.Semester
	}
	if updateUserData.IPK != 0.0 {
		user.IPK = updateUserData.IPK
	}

	// Save the Changes
	db.Save(&user)
	// Return the updated user
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "users Found", "data": user})
}

// delete user in db by ID
func DeleteUserByID(c *fiber.Ctx) error {
	db := database.DB.Db
	var user model.User
	// get id dengan parameter
	id := c.Params("id")

	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "User not found", "data": nil})
	}

	//hapus user, jika error maka err, jika tidak return status 200
	err := db.Delete(&user, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete user", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "User deleted"})
}
