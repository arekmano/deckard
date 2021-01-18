package cmd

import (
	"fmt"

	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/reporter"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func executeCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "execute",
		Short: "execute",
		Long:  `m`,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logrus.New()
			logger.SetLevel(logrus.DebugLevel)
			e := executor.New(&reporter.PrintReporter{Logger: logger}, logrus.NewEntry(logger), func(context interface{}) error {
				fmt.Println("EXECUTING")
				return nil
			})
			e.Execute(nil)
			return nil
		},
	}
	return command
}
