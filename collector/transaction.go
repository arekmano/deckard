package collector

type TransactionContext string

type Transaction = func(context TransactionContext) (interface{}, error)

type TransactionConfig struct{}
