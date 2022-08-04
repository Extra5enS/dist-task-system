package taskBuilder

import (
	"fmt"
	"strings"

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

func (ite ExtTaskExecutor) Exec(t Task) (string, error) {
	if ite.oph == nil {
		return "", fmt.Errorf("unvalid ouptuter")
	}
	if exfun, ok := ExtTaskExecutionTable[t.Name]; !ok {
		return "", fmt.Errorf("Unknowen command name")
	} else {
		return exfun(ite.oph, t)
	}
}

var ExtTaskExecutionTable = map[string](func(o outputer.Outputer, t Task) (string, error)){
	"foreach": func(o outputer.Outputer, t Task) (string, error) {
		out := o.Get(t.Args[0], t.Args[1:], t.IncomeAddr)
		anss := make([]string, 0)
		for i := 0; i < o.AnsCount(); i++ {
			anss = append(anss, <-out)
		}
		return strings.Join(anss, " "), nil
	},
}
