package stress

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/reporter"
	"github.com/arekmano/deckard/stats"
	"github.com/arekmano/deckard/transaction"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func fixedTpsCommand() *cobra.Command {
	var binarypath *string
	var binaryargs *[]string
	var durationArg *string
	var tpsArg *float64
	var reportIntervalArg *float64
	command := &cobra.Command{
		Use:   "fixed-tps",
		Short: "fixed-tps",
		Long:  `m`,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logrus.New()
			e := executor.New(logrus.NewEntry(logger), transaction.Binary(*binarypath, *binaryargs))
			interval := time.Duration(float64(time.Second) * 1 / *tpsArg)
			var stopTime time.Time
			if *durationArg == "" {
				stopTime = time.Now().Add(time.Hour * 999999)
			} else {
				duration, err := time.ParseDuration(*durationArg)
				if err != nil {
					return err
				}
				stopTime = time.Now().Add(duration)
			}
			var startedTransactions int64
			var reports []executor.Report
			startTime := time.Now()
			var starts []time.Time
			// reporter
			go func() {
				for {
					reportInterval := time.Duration(float64(time.Second) * *reportIntervalArg)
					time.Sleep(reportInterval)
					currTime := time.Now()
					stat, err := stats.StatsBetween(reports, starts, currTime.Add(-1*reportInterval), currTime)
					if err != nil {
						logger.WithError(err).Error("could not calculate stats")
						continue
					}
					logger.
						WithField("Running", len(reports)-int(startedTransactions)).
						WithField("Ended", len(reports)).
						WithField("TPS (estimate)", float64(stat.StartedTransactionCount)/5).
						WithField(fmt.Sprintf("Failed (Last %.2f seconds)", reportInterval.Seconds()), stat.FailedTransactionCount).
						WithField(fmt.Sprintf("Success (Last %.2f seconds)", reportInterval.Seconds()), stat.SuccessfulTransactionCount).
						WithField(fmt.Sprintf("P99 (ms - Last %.2f seconds)", reportInterval.Seconds()), stat.P99ExecutionTime).
						WithField(fmt.Sprintf("P90 (ms - Last %.2f seconds)", reportInterval.Seconds()), stat.P90ExecutionTime).
						WithField(fmt.Sprintf("Mean (ms - Last %.2f seconds)", reportInterval.Seconds()), stat.MeanExecutionTime).
						WithField("Total runtime", currTime.Sub(startTime).Seconds()).
						Info()
				}
			}()
			// Metric collector
			reportChan := make(chan executor.Report)
			go func() {
				for {
					select {
					case report := <-reportChan:
						reports = append(reports, report)
					}
				}
			}()

			for time.Now().Before(stopTime) {
				atomic.AddInt64(&startedTransactions, 1)
				starts = append(starts, time.Now())
				go func() {
					report, _ := e.Execute(nil)
					reportChan <- *report
				}()
				time.Sleep(interval)
			}
			for {
				if startedTransactions == int64(len(reports)) {
					v := reporter.PrintReporter{Logger: logger}
					s, err := stats.Stats(reports, starts)
					if err != nil {
						return err
					}
					v.ReportStatistics(s)
					return nil
				}
				logger.
					WithField("Running", startedTransactions-int64(len(reports))).
					Info("Waiting on executing transactions to complete")
				time.Sleep(time.Second)
			}
		},
	}
	binarypath = command.Flags().StringP("binarypath", "b", "", "the path to the binary to execute")
	binaryargs = command.Flags().StringArrayP("args", "a", []string{}, "args")
	durationArg = command.Flags().StringP("duration", "d", "", "the path to the binary to execute")
	reportIntervalArg = command.Flags().Float64P("reportInterval", "r", 5, "the interval with which to report stats")
	tpsArg = command.Flags().Float64P("tps", "t", 0, "the path to the binary to execute")
	command.MarkFlagRequired("binarypath")
	command.MarkFlagRequired("tps")
	return command
}
