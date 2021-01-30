package reporter

import (
	"io"

	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/stats"
)

type Reporter interface {
	ReportTransaction(r *executor.Report) error
	ReportStatistics(s *stats.Statistics) error
	Writer() io.Writer
}
