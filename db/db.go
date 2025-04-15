package db

import (
	"log"
	"os"
	"savegen-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
}

func Connect() error {
	dbConfig, err := config.NewDBConfig()
	if err != nil {
		return err
	}

	db, err = gorm.Open(postgres.Open(dbConfig.DSN()), &gorm.Config{
		Logger: getLogger(),
	})
	if err != nil {
		return err
	}
	
	log.Println("Database connected successfully")
	return nil
}

func Get() *gorm.DB {
	return db
}

func Close() error {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		return sqlDB.Close()
	}
	return nil
}