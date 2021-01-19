package reporter

import (
	"time"

	"github.com/arekmano/deckard/transaction"
)

type Reporter interface {
	Report(r *Report) error
}

type Report struct {
	Status    transaction.TransactionStatus
	StartTime time.Time
	EndTime   time.Time
	Message   string
}
