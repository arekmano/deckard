package reporter

import "github.com/arekmano/deckard/executor"

type DynamoDBReporter struct {
}

func (c *DynamoDBReporter) ReportTransaction(r *executor.Report) error {
	return nil
}
