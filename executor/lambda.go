package executor

type LambdaExecutor struct {
}

func (e *Executor) execute(t Transaction) {
	res, err := t("c")
	if err != nil {

	}
}
