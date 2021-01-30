package stress

import (
	"time"

	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/reporter"
	"github.com/arekmano/deckard/service"
	"github.com/arekmano/deckard/transaction"
	"github.com/pkg/errors"
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
			stopTime, err := ParseStopTime(*durationArg)
			if err != nil {
				return err
			}
			d := service.New(&reporter.PrintReporter{Logger: logger}, e, logrus.NewEntry(logger).WithField("mode", "fixed-tps"), *stopTime, interval)
			return d.FixedTps(*reportIntervalArg)
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

func ParseStopTime(durationArg string) (stopTime *time.Time, err error) {
	var t time.Time
	if durationArg == "" {
		t = time.Now().Add(time.Hour * 999999)
	} else {
		duration, err := time.ParseDuration(durationArg)
		if err != nil {
			return nil, errors.Wrap(err, "duration argument invalid")
		}
		t = time.Now().Add(duration)
	}
	return &t, nil
}
