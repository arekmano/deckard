package transaction

import (
	"os/exec"
)

func Binary(binaryPath string, binaryargs []string) Transaction {
	return func(context *TransactionContext) error {
		command := exec.Command(binaryPath, binaryargs...)
		command.Stdout = context.TransactionWriter
		command.Stderr = context.TransactionWriter
		return command.Run()
	}
}
