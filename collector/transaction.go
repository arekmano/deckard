package collector

const (
	Fail    TransactionStatus = "Fail"
	Success TransactionStatus = "Success"
	Timeout TransactionStatus = "Timeout"
)

type TransactionContext string
type TransactionStatus string
type Transaction = func(context TransactionContext) (interface{}, error)

type TransactionConfig struct {
}
