package database

import (
	"github.com/arturbaccarin/credit-card-messaging/register-service/internal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func StartDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("../../internal/database/registerDB.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&model.AuditOrder{}, &model.PaymentOrder{})
	if err != nil {
		panic(err)
	}

	return db
}
