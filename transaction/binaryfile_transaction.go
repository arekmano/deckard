package transaction

import (
	"os/exec"
)

func Binary(binaryPath string, binaryargs []string) Transaction {
	return func(context interface{}) error {
		command := exec.Command(binaryPath, binaryargs...)
		// command.Stdout = os.Stdout
		// command.Stderr = os.Stderr
		return command.Run()
	}
}
