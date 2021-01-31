package stress

import (
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func StressCommand() *cobra.Command {
	stressCmd := &cobra.Command{
		Use:   "stress",
		Short: "Executes a stress test",
		Long:  `Executes a stress test using the given tool`,
	}

	stressCmd.AddCommand(fixedTpsCommand(), maxParallelCommand())
	return stressCmd
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
