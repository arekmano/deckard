package reporter

import "github.com/arekmano/deckard/executor"

type NullReporter struct {
}

func (c *NullReporter) ReportTransaction(r *executor.Report) error {
	return nil
}
