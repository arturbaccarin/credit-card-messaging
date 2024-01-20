package model

import (
	"time"

	"github.com/google/uuid"
)

type AuditOrder struct {
	ID             int
	paymentOrderID uuid.UUID
	Date           time.Time
}
