package taskBuilder

import (
	"fmt"
	"strconv"
)

type IntTaskExecutor struct {
}

func (ite IntTaskExecutor) Exec(t Task) (string, error) {
	if exfun, ok := IntTaskExecutionTable[t.Name]; !ok {
		return "", fmt.Errorf("Unknowen command name")
	} else {
		return exfun(t)
	}
}

var IntTaskExecutionTable = map[string](func(t Task) (string, error)){
	"hello": func(t Task) (string, error) {
		return "hello, master!", nil
	},
	"sum": func(t Task) (string, error) {
		sum := 0.0
		for i, arg := range t.Args {
			val, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				return "", fmt.Errorf("Arg[%d] have problem: %v", i, err)
			}
			sum += val
		}
		return fmt.Sprintf("%v", sum), nil
	},
	"nothing": func(t Task) (string, error) {
		return "", nil
	},
}
