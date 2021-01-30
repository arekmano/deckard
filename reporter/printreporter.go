package reporter

import (
	"io"

	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/stats"
	"github.com/sirupsen/logrus"
)

type PrintReporter struct {
	Logger *logrus.Logger
}

func (c *PrintReporter) ReportTransaction(r *executor.Report) error {
	c.Logger.
		WithField("StartTime", r.StartTime).
		WithField("EndTime", r.EndTime).
		WithField("Duration", r.EndTime.Sub(r.StartTime)).
		WithField("Status", r.Status).
		WithField("Message", r.Message).
		Info("Transaction Report")
	return nil
}

func (c *PrintReporter) ReportStatistics(stat *stats.Statistics) error {
	runtimeSeconds := stat.EndTime.Sub(stat.StartTime).Seconds()
	c.Logger.
		WithField("Running", stat.TotalTransactionCount-stat.TotalStartedTransactionCount).
		WithField("Ended", stat.TotalEndedTransactionCount).
		WithField("TPS", float64(stat.TotalTransactionCount)/runtimeSeconds).
		WithField("Failed", stat.TotalFailedTransactionCount).
		WithField("Success", stat.SuccessfulTransactionCount).
		WithField("Overall TPS", float64(stat.TotalTransactionCount)/runtimeSeconds).
		WithField("P99 (ms)", stat.P99ExecutionTime).
		WithField("P90 (ms)", stat.P90ExecutionTime).
		WithField("Mean (ms)", stat.MeanExecutionTime).
		WithField("Runtime", runtimeSeconds).
		WithField("Start time", stat.StartTime).
		WithField("End time", stat.EndTime).
		Info("Statistic Report")
	return nil
}

func (c *PrintReporter) Writer() io.Writer {
	return c.Logger.Out
}
