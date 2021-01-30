package executor

import (
	"time"

	"github.com/arekmano/deckard/transaction"
	"github.com/sirupsen/logrus"
)

type Executor struct {
	logger          *logrus.Entry
	transactionFunc transaction.Transaction
}

type Report struct {
	Status    transaction.TransactionStatus
	StartTime time.Time
	EndTime   time.Time
	Message   string
}

func New(logger *logrus.Entry, transactionFunc transaction.Transaction) *Executor {
	return &Executor{
		logger:          logger,
		transactionFunc: transactionFunc,
	}
}

func (e *Executor) Execute(context *transaction.TransactionContext) (*Report, error) {
	e.logger.
		WithField("Status", transaction.Initializing).
		Debug("Starting transaction")
	startTime := time.Now()
	err := e.transactionFunc(context)
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
	report := &Report{
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

	return report, nil
}
