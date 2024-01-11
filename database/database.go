package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/Arrasty/tugas/config"
	"github.com/Arrasty/tugas/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// struct instance
type Dbinstance struct {
	Db *gorm.DB //
}

var DB Dbinstance

// Connect function
func Connect() {
	p := config.Config("DB_PORT")

	//karena return dari config functionnya adalah string
	// dilakukan parsing string ke unsigned int menggunakan paket strconv dengan fungsi ParseUint
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		fmt.Println("Error parsing string to int")
	}

	//buat dsn untuk dapat terkoneksi dengan database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Config("DB_HOST"), config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME"), port)

	//connect ke daatabase
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Gagal terkoneksi dengan database. \n", err)
		os.Exit(2)
	}

	//migrasi otomatis ke model user dan menyimpan database di variable DB
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&model.User{})
	DB = Dbinstance{
		Db: db,
	}
}
