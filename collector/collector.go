package collector

type Executor struct{}

func (e *Executor) execute(t Transaction) {
	t("c")
}
