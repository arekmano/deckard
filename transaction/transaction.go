package transaction

import "io"

const (
	Fail         TransactionStatus = "Fail"
	Success      TransactionStatus = "Success"
	Initializing TransactionStatus = "Initializing"
	Timeout      TransactionStatus = "Timeout"
)

type TransactionStatus string
type Transaction = func(context *TransactionContext) error

type TransactionContext struct {
	TransactionWriter io.Writer
}
