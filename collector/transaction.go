package collector

const (
	Fail         TransactionStatus = "Fail"
	Success      TransactionStatus = "Success"
	Initializing TransactionStatus = "Initializing"
	Timeout      TransactionStatus = "Timeout"
)

type TransactionStatus string
type Transaction = func(context interface{}) error

type TransactionConfig struct {
}
