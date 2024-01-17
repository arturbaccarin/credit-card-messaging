package model

import (
	"time"

	"github.com/google/uuid"
)

type PaymentOrder struct {
	ID        uuid.UUID
	Value     float64
	Date      time.Time
	NeedAudit bool
}
