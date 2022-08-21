package db

import (
	"github.com/bagasss3/go-weight/internal/config"
	"github.com/bagasss3/go-weight/internal/model"

	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	PostgreDB *gorm.DB
)

func InitializeDbConn() {
	dialect := postgres.Open(config.DatabaseDSN())
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		logger.WithField("databaseDSN", config.DatabaseDSN()).Fatal("failed to connect database: ", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		logger.WithField("databaseDSN", config.DatabaseDSN()).Fatal("failed to connect database: ", err)
	}
	sqlDB.SetMaxIdleConns(config.PostgreMaxIdleConns())
	sqlDB.SetMaxOpenConns(config.PostgreMaxOpenConns())
	sqlDB.SetConnMaxLifetime(config.PostgreSetConnMaxLifeTime())

	PostgreDB = db

	logger.Info("Connection to postgreql success...")

}

func AutoMigrate() error {
	if !PostgreDB.Migrator().HasTable(&model.Weight{}) {
		err := PostgreDB.AutoMigrate(&model.Weight{})
		return err
	}
	return nil
}
