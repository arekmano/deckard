package reporter

import (
	"time"

	"github.com/arekmano/deckard/collector"
)

type Reporter interface {
	Report(r *Report) error
}

type Report struct {
	status    collector.TransactionStatus
	startTime time.Time
	endTime   time.Time
	message   string
}
