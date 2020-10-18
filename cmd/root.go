package cmd

import "github.com/spf13/cobra"

func GetCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   "deckard",
		Short: "deckard",
		Long:  `m`,
	}

	root.AddCommand(echoCommand())
	return root
}
