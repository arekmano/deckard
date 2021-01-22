package cmd

import (
	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/transaction"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func singleCommand() *cobra.Command {
	var binarypath *string
	var binaryargs *[]string
	command := &cobra.Command{
		Use:   "single",
		Short: "Executes the tool a single time",
		Long:  `m`,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logrus.New()
			logger.SetLevel(logrus.DebugLevel)
			e := executor.New(logrus.NewEntry(logger), transaction.Binary(*binarypath, *binaryargs))
			e.Execute(nil)
			return nil
		},
	}
	binarypath = command.Flags().StringP("binarypath", "b", "", "the path to the binary to execute")
	binaryargs = command.Flags().StringArrayP("args", "a", []string{}, "args")
	command.MarkFlagRequired("binarypath")

	return command
}
