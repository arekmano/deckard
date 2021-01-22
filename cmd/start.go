package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func startCommand() *cobra.Command {
	var token *string
	command := &cobra.Command{
		Use:   "start",
		Short: "Starts the long-running periodic execution of some command",
		Long:  `Starts the long-running periodic execution of some command. Will persist until "deckard stop" is called.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("token: %s\n", *token)
			return nil
		},
	}

	token = command.Flags().StringP("token", "t", "", "test token")
	command.MarkFlagRequired("token")

	return command
}
