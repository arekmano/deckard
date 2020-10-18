package executor

import (
	"time"

	"github.com/arekmano/deckard/collector"
	"github.com/arekmano/deckard/reporter"
	"github.com/sirupsen/logrus"
)

type Executor struct {
	reporter reporter.Reporter
	logger   *logrus.Entry
}

func (e *Executor) execute(t collector.Transaction) {
	startTime := time.Now()
	res, err := t("c")
	e.logger.
		WithField("result", res).
		Info("Transaction Completed")
	endTime := time.Now()
	var status collector.TransactionStatus
	var message string
	if err != nil {
		status = collector.Fail
		message = err.Error()
	} else {
		status = collector.Success
	}
	report := &reporter.Report{
		startTime: startTime,
		endTime:   endTime,
		status:    status,
		message:   message,
	}
	if err != nil {
		e.reporter.Report(report)
	}
}
