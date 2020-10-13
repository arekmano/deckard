package collector

import "github.com/arekmano/deckard/reporter"

type Executor struct {
	reporter reporter.Reporter
}

func (e *Executor) execute(t Transaction) {
	res, err := t("c")
	if err != nil {
		e.reporter()
	}
}
