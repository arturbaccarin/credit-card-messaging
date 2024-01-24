package model

import (
	"time"

	"github.com/google/uuid"
)

type PaymentOrder struct {
	ID        uuid.UUID `json:"id"`
	Value     float64   `json:"value"`
	Date      time.Time `json:"date"`
	NeedAudit bool      `json:"need_audit"`
}
