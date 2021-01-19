package executor

import (
	"time"

	"github.com/arekmano/deckard/reporter"
	"github.com/arekmano/deckard/transaction"
	"github.com/sirupsen/logrus"
)

type ExecutorService struct {
	reporter        reporter.Reporter
	logger          *logrus.Entry
	transactionFunc transaction.Transaction
}

func New(report reporter.Reporter, logger *logrus.Entry, transactionFunc transaction.Transaction) *ExecutorService {
	return &ExecutorService{
		reporter:        report,
		logger:          logger,
		transactionFunc: transactionFunc,
	}
}

func (e *ExecutorService) Execute(input interface{}) {
	e.logger.
		WithField("Status", transaction.Initializing).
		Debug("Starting transaction")
	startTime := time.Now()
	err := e.transactionFunc(input)
	endTime := time.Now()
	var status transaction.TransactionStatus
	var message string
	if err != nil {
		status = transaction.Fail
		message = err.Error()
	} else {
		status = transaction.Success
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
