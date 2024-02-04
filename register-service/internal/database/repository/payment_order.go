package repository

import (
	"time"

	"github.com/arturbaccarin/credit-card-messaging/register-service/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentOrder struct {
	db *gorm.DB
}

func NewPaymentOrder(db *gorm.DB) *PaymentOrder {
	return &PaymentOrder{
		db: db,
	}
}

func (p PaymentOrder) Create(paymentOrderID uuid.UUID, value float64, date time.Time) error {
	paymentOrder := model.PaymentOrder{
		PaymentOrderID: paymentOrderID,
		Value:          value,
		Date:           date,
	}

	result := p.db.Create(&paymentOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
