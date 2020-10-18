package executor

import "github.com/arekmano/deckard/collector"

type LambdaExecutor struct {
}

func (e *LambdaExecutor) execute(t collector.Transaction) {
	res, err := t("c")
	if err != nil {
		return
	}
}
