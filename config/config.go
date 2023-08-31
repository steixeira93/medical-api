package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *logger
)

func Init() error {
	var err error
	db, err := InitializeSqlite()

	if err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	}
	return nil
}

func GetSQLiteDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}