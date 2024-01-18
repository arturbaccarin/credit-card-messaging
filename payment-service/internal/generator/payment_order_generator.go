package generator

import (
	"math/rand"
	"time"

	"github.com/arturbaccarin/credit-card-messaging/payment-service/internal/model"
	"github.com/google/uuid"
)

type PaymentOrderGenerator struct {
	mumberOfCopies int
}

func NewPaymentOrderGenerator(numberOfCopies int) *PaymentOrderGenerator {
	return &PaymentOrderGenerator{
		mumberOfCopies: numberOfCopies,
	}
}

func (s PaymentOrderGenerator) GenerateCopies() []model.PaymentOrder {
	paymentOrders := make([]model.PaymentOrder, 0, s.mumberOfCopies)

	for i := 0; i < s.mumberOfCopies; i++ {
		paymentOrder := model.PaymentOrder{
			ID:        uuid.New(),
			Value:     rand.Float64() * 100.0,
			Date:      time.Now(),
			NeedAudit: s.generateRandomNeedAudit(),
		}

		paymentOrders = append(paymentOrders, paymentOrder)
	}

	return paymentOrders
}

func (PaymentOrderGenerator) generateRandomNeedAudit() bool {
	randomNumber := rand.Intn(2)

	return randomNumber == 1
}
