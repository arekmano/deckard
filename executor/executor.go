package executor

import (
	"time"

	"github.com/arekmano/deckard/collector"
	"github.com/arekmano/deckard/reporter"
	"github.com/sirupsen/logrus"
)

type ExecutorService struct {
	reporter        reporter.Reporter
	logger          *logrus.Entry
	transactionFunc collector.Transaction
}

func New(report reporter.Reporter, logger *logrus.Entry, transactionFunc collector.Transaction) *ExecutorService {
	return &ExecutorService{
		reporter:        report,
		logger:          logger,
		transactionFunc: transactionFunc,
	}
}

func (e *ExecutorService) Execute(input interface{}) {
	e.logger.
		WithField("Status", collector.Initializing).
		Debug("Starting transaction")
	startTime := time.Now()
	err := e.transactionFunc(input)
	endTime := time.Now()
	var status collector.TransactionStatus
	var message string
	if err != nil {
		status = collector.Fail
		message = err.Error()
	} else {
		status = collector.Success
		message = "Completed Successfully"
	}
	report := &reporter.Report{
		StartTime: startTime,
		EndTime:   endTime,
		Status:    status,
		Message:   message,
	}
	e.logger.
		WithField("StartTime", startTime).
		WithField("EndTime", endTime).
		WithField("Status", status).
		WithField("Message", message).
		Debug("Sending report")

	e.reporter.Report(report)
	return
}
