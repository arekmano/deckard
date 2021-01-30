package executor_test

import (
	"fmt"
	"testing"

	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/transaction"
	"github.com/sirupsen/logrus"
)

var testFunction = func(context *transaction.TransactionContext) error {
	fmt.Println("FUNCTION EXECUTING")
	return nil
}

func TestExecute_WithSuccess(t *testing.T) {
	logger := logrus.New()
	e := executor.New(logrus.NewEntry(logger), testFunction)
	e.Execute(nil)
}
