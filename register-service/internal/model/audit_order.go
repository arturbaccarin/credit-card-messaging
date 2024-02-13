package model

import (
	"time"

	"github.com/google/uuid"
)

type AuditOrder struct {
	ID             int
	PaymentOrderID uuid.UUID
	Date           time.Time
}
