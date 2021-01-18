package reporter

import (
	"time"

	"github.com/arekmano/deckard/collector"
)

type Reporter interface {
	Report(r *Report) error
}

type Report struct {
	Status    collector.TransactionStatus
	StartTime time.Time
	EndTime   time.Time
	Message   string
}
