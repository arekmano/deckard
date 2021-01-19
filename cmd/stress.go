package cmd

import (
	"math"

	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/reporter"
	"github.com/arekmano/deckard/transaction"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func stressCommand() *cobra.Command {
	var binarypath *string
	var binaryargs *[]string
	var frequency *int
	var maxParallel *int
	command := &cobra.Command{
		Use:   "stress",
		Short: "stress",
		Long:  `m`,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logrus.New()
			logger.SetLevel(logrus.DebugLevel)
			e := executor.New(&reporter.PrintReporter{Logger: logger}, logrus.NewEntry(logger), transaction.Binary(*binarypath, *binaryargs))
			c := make(chan int, *maxParallel)
			initialThreads := int(math.Min(float64(*frequency), float64(*maxParallel)))
			invocationsComplete := 0
			for i := 0; i < initialThreads; i++ {
				logger.Debug("Loading Up initial thread")
				go func(threadNumber int) {
					e.Execute(nil)
					c <- threadNumber
				}(i)
			}
			for {
				select {
				case i := <-c:
					invocationsComplete++
					if invocationsComplete < *frequency {
						logger.
							WithField("InvocationsLeft", *frequency-invocationsComplete).
							WithField("InvocationsComplete", invocationsComplete).
							WithField("Thread", i).
							Debug()
						go func(threadNumber int) {
							e.Execute(nil)
							c <- threadNumber
						}(i)
					} else if invocationsComplete == *frequency {
						logger.Debug("Did all my work. are you proud?")
						return nil
					}
				}
			}
		},
	}
	binarypath = command.Flags().StringP("binarypath", "b", "", "the path to the binary to execute")
	binaryargs = command.Flags().StringArrayP("args", "a", []string{}, "args")
	frequency = command.Flags().IntP("frequency", "f", 1, "the number of executions to perform")
	maxParallel = command.Flags().IntP("threads", "t", 1, "the interval to wait between executions")
	command.MarkFlagRequired("binarypath")

	return command
}
