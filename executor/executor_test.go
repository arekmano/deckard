package executor_test

import (
	"fmt"
	"testing"

	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/reporter"
	"github.com/sirupsen/logrus"
)

var testFunction = func(context interface{}) error {
	fmt.Println("FUNCTION EXECUTING")
	return nil
}

func TestExecute_WithSuccess(t *testing.T) {
	logger := logrus.New()
	e := executor.New(&reporter.PrintReporter{Logger: logger}, logrus.NewEntry(logger), testFunction)
	e.Execute(nil)
}
