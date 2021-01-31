package stress

import (
	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/reporter"
	"github.com/arekmano/deckard/service"
	"github.com/arekmano/deckard/transaction"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func maxParallelCommand() *cobra.Command {
	var binarypath *string
	var binaryargs *[]string
	var durationArg *string
	var reportIntervalArg *float64
	var maxParallel *int
	command := &cobra.Command{
		Use:   "max-parallel",
		Short: "max-parallel",
		Long:  `m`,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logrus.New()
			stopTime, err := ParseStopTime(*durationArg)
			if err != nil {
				return err
			}
			e := executor.New(logrus.NewEntry(logger), transaction.Binary(*binarypath, *binaryargs))
			d := service.New(&reporter.PrintReporter{Logger: logger}, e, logrus.NewEntry(logger).WithField("mode", "fixed-tps"), *stopTime)
			return d.MaxParallel(*maxParallel, *reportIntervalArg)

		},
	}
	binarypath = command.Flags().StringP("binarypath", "b", "", "the path to the binary to execute")
	binaryargs = command.Flags().StringArrayP("args", "a", []string{}, "args")
	durationArg = command.Flags().StringP("duration", "d", "", "the path to the binary to execute")
	reportIntervalArg = command.Flags().Float64P("reportInterval", "r", 5, "the interval with which to report stats")
	maxParallel = command.Flags().IntP("threads", "t", 1, "the interval to wait between executions")
	command.MarkFlagRequired("binarypath")

	return command
}
