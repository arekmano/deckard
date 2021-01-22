package service

import (
	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/reporter"
	"github.com/sirupsen/logrus"
)

type Deckard struct {
	reporter reporter.Reporter
	executor executor.Executor
	logger   *logrus.Entry
}
