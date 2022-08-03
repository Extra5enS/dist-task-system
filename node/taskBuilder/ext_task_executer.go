package taskBuilder

import (
	"fmt"

	"github.com/Extra5enS/dist-task-system/node/outputer"
)

type ExtTaskExecutor struct {
	oph outputer.Outputer
}

func NewExtTaskExecutor(oph outputer.Outputer) ExtTaskExecutor {
	return ExtTaskExecutor{
		oph: oph,
	}
}

func (ite ExtTaskExecutor) Exec(task Task) ([]string, error) {
	if exfun, ok := ExtTaskExecutionTable[task.Name]; !ok {
		return []string{}, fmt.Errorf("Unknowen command name")
	} else {
		return exfun(ite.oph, task)
	}
}

var ExtTaskExecutionTable = map[string](func(o outputer.Outputer, task Task) ([]string, error)){
	"forevery": func(o outputer.Outputer, task Task) ([]string, error) {
		o.Get(task.Name, task.Args)
		return []string{}, nil
	},
}
