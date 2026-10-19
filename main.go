package main

import (
	// "fmt"
	// "ginframework/database"
	// "ginframework/router"
	// "log"

	"fmt"
	"log"

	"ginframework/database"
	"ginframework/router"
	"ginframework/todo"
)

// ================ database/sql Method ================
// func main() {
// 	db, err := database.Connect()
// 	if err != nil {
// 		log.Fatal("Gagal terhubung ke database: ", err)
// 	}
// 	defer db.Close()
// 	fmt.Println("✅ Terhubung ke database PostgreSQL")

// 	r := router.SetupRouter(db)

// 	fmt.Println("🚀 Server Gin berjalan di http://localhost:8080")
// 	if err := r.Run(":8080"); err != nil {
// 		log.Fatal("Gagal menjalankan server: ", err)
// 	}
// }

// ================ GORM Method ================
func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Gagal terhubung ke database: ", err)
	}
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	fmt.Println("✅ Terhubung ke database PostgreSQL")

	// AutoMigrate: membuat tabel todos jika belum ada, atau menambah kolom yang kurang
	if err := db.AutoMigrate(&todo.Todo{}); err != nil {
		log.Fatal("Gagal migrasi database: ", err)
	}

	r := router.SetupRouter(db)

	fmt.Println("🚀 Server Gin berjalan di http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Gagal menjalankan server: ", err)
	}
}
