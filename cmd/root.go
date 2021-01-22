package cmd

import (
	"github.com/arekmano/deckard/cmd/stress"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "deckard",
		Short: "deckard",
		Long:  `m`,
	}

	root.AddCommand(echoCommand(), stress.StressCommand(), singleCommand(), setupCommand())
	return root
}
