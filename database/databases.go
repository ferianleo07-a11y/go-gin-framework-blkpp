package database

import (
	"context"
	"errors"
	"os"
	"time"

	// _ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// ================ database/sql Method ================
// func Connect() (*sql.DB, error) {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	dsn := os.Getenv("DATABASE_URL")
// 	if dsn == "" {
// 		return nil, errors.New("environment variable DATABASE_URL belum diatur")
// 	}

// 	var err error
// 	db, err := sql.Open("pgx", dsn)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
// 	defer cancel()

// 	if err := db.PingContext(ctx); err != nil {
// 		db.Close()
// 		return nil, err
// 	}

// 	return db, nil
// }

// ================ GORM Method ================
func Connect() (*gorm.DB, error) {
	// Load file .env
	_ = godotenv.Load()

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		return nil, errors.New("environment variable DATABASE_URL belum diatur")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Info = tampilkan setiap SQL yang dibuat GORM (bagus untuk belajar).
		// Ganti ke logger.Warn jika terminal terlalu ramai.
		Logger:               logger.Default.LogMode(logger.Info),
		DisableAutomaticPing: true, // ping dilakukan manual di bawah, dengan batas waktu
	})
	if err != nil {
		return nil, err
	}

	// Ping untuk memastikan database benar-benar bisa dihubungi
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		sqlDB.Close()
		return nil, err
	}

	return db, nil
}
