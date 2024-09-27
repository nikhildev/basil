package invoicer

import (
	"math/rand"
	"time"
)

type Invoice struct {
	ID        int64
	Amount    float64
	CreatedAt time.Time
}

func GetNewInvoice() *Invoice {
	return &Invoice{
		ID:        int64(rand.Int()),
		Amount:    rand.Float64(),
		CreatedAt: time.Now(),
	}
}
