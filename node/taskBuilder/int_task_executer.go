package taskBuilder

import (
	"fmt"
	"strconv"
)

type IntTaskExecutor struct {
}

func (ite IntTaskExecutor) Exec(name string, args []string) (string, error) {
	if exfun, ok := IntTaskExecutionTable[name]; !ok {
		return "", fmt.Errorf("Unknowen command name")
	} else {
		return exfun(args)
	}
}

var IntTaskExecutionTable = map[string](func(args []string) (string, error)){
	"hello": func(args []string) (string, error) {
		return "hello, master!", nil
	},
	"sum": func(args []string) (string, error) {
		sum := 0.0
		for i, arg := range args {
			val, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				return "", fmt.Errorf("Arg[%d] have problem: %v", i, err)
			}
			sum += val
		}
		return fmt.Sprintf("%v", sum), nil
	},
}
