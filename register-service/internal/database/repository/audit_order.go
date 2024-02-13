package repository

import (
	"time"

	"github.com/arturbaccarin/credit-card-messaging/register-service/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AuditOrder struct {
	db *gorm.DB
}

func NewAuditOrder(db *gorm.DB) *AuditOrder {
	return &AuditOrder{
		db: db,
	}
}

func (a AuditOrder) Create(paymentOrderID uuid.UUID, date time.Time) error {
	auditOrder := model.AuditOrder{
		PaymentOrderID: paymentOrderID,
		Date:           date,
	}

	result := a.db.Create(&auditOrder)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
