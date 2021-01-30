package transaction

func GolangTransaction(binaryPath string, binaryargs []string) Transaction {
	return func(context *TransactionContext) error {
		return nil
	}
}
