package service

import (
	"time"

	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/reporter"
	"github.com/arekmano/deckard/stats"
	"github.com/arekmano/deckard/transaction"
	"github.com/sirupsen/logrus"
)

type Deckard struct {
	reporter      reporter.Reporter
	executor      *executor.Executor
	logger        *logrus.Entry
	reports       []executor.Report
	startTime     time.Time
	startInterval time.Duration
	stopTime      time.Time
	startTimes    []time.Time
	reportChan    chan executor.Report
}

func New(r reporter.Reporter, e *executor.Executor, l *logrus.Entry, stopTime time.Time, interval time.Duration) *Deckard {
	return &Deckard{
		reporter:      r,
		executor:      e,
		logger:        l,
		reports:       []executor.Report{},
		startTime:     time.Now(),
		startInterval: interval,
		stopTime:      stopTime,
		startTimes:    []time.Time{},
		reportChan:    make(chan executor.Report),
	}
}

func (d *Deckard) FixedTps(reportIntervalArg float64) error {
	go d.reporterRoutine(reportIntervalArg)
	go d.metricCollectorRoutine()
	d.executionStarter()
	for {
		if len(d.startTimes) == len(d.reports) {
			s, err := stats.Stats(d.reports, d.startTimes)
			if err != nil {
				return err
			}
			d.reporter.ReportStatistics(s)
			return nil
		}
		d.logger.
			WithField("Running", len(d.startTimes)-len(d.reports)).
			Info("Waiting on executing transactions to complete")
		time.Sleep(time.Second)
	}
}

func (d *Deckard) reporterRoutine(reportIntervalArg float64) {
	for {
		reportInterval := time.Duration(float64(time.Second) * reportIntervalArg)
		time.Sleep(reportInterval)
		currTime := time.Now()
		stat, err := stats.StatsBetween(d.reports, d.startTimes, currTime.Add(-1*reportInterval), currTime)
		if err != nil {
			d.logger.WithError(err).Error("could not calculate stats")
			continue
		}
		d.reporter.ReportStatistics(stat)
	}
}

func (d *Deckard) executionStarter() {
	for time.Now().Before(d.stopTime) {
		d.startTimes = append(d.startTimes, time.Now())
		go func() {
			report, _ := d.executor.Execute(&transaction.TransactionContext{
				TransactionWriter: d.reporter.Writer(),
			})
			d.reportChan <- *report
		}()
		time.Sleep(d.startInterval)
	}
}

func (d *Deckard) metricCollectorRoutine() {
	for {
		select {
		case report := <-d.reportChan:
			d.reports = append(d.reports, report)
			d.reporter.ReportTransaction(&report)
		}
	}
}
