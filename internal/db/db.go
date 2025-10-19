package db

import (
	"fmt"
	// "time"

	"github.com/MiroslavArbor/student-organisations-api/internal/config"
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"github.com/MiroslavArbor/student-organisations-api/internal/seeding"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func createConnection(config *config.Config) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB(config *config.Config) (*gorm.DB, error) {
	db, err := createConnection(config)
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)

	err = db.AutoMigrate(
		&models.Organisation{},
		&models.Role{},
		&models.Member{},
		&models.UserRole{},
		&models.Team{},
		&models.TeamMember{},
		&models.Project{},
		&models.MemberProject{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SeedDatabase(db *gorm.DB, config *config.Config) error {
	if config.ENV == "development" {
		return seeding.InsertTestData(db)
	}
	return nil
}
