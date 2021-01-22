package stress

import "github.com/spf13/cobra"

func StressCommand() *cobra.Command {
	stressCmd := &cobra.Command{
		Use:   "stress",
		Short: "Executes a stress test",
		Long:  `Executes a stress test using the given tool`,
	}

	stressCmd.AddCommand(fixedTpsCommand(), maxParallelCommand())
	return stressCmd
}
