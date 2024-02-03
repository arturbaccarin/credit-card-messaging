package model

import (
	"time"

	"github.com/google/uuid"
)

type PaymentOrder struct {
	ID             int
	PaymentOrderID uuid.UUID
	Value          float64
	Date           time.Time
}
