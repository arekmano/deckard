package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func echoCommand() *cobra.Command {
	var token *string
	command := &cobra.Command{
		Use:   "echo",
		Short: "echo",
		Long:  `m`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("token: %s\n", *token)
			return nil
		},
	}

	token = command.Flags().StringP("token", "t", "", "test token")
	command.MarkFlagRequired("token")

	return command
}
