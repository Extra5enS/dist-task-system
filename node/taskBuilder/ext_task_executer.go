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

func (ite ExtTaskExecutor) Exec(name string, args []string) (string, error) {
	if ite.oph == nil {
		return "", fmt.Errorf("unvalid ouptuter")
	}
	if exfun, ok := ExtTaskExecutionTable[name]; !ok {
		return "", fmt.Errorf("Unknowen command name")
	} else {
		return exfun(ite.oph, args)
	}
}

var ExtTaskExecutionTable = map[string](func(o outputer.Outputer, args []string) (string, error)){
	"forevery": func(o outputer.Outputer, args []string) (string, error) {
		out := o.Get(args[0], args[1:])
		anss := make([]string, 0)
		for i := 0; i < o.AnsCount(); i++ {
			anss = append(anss, <-out)
		}
		return strings.Join(anss, " "), nil
	},
}
