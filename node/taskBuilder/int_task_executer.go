package taskBuilder

import "fmt"

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
		return "hello, master!\n", nil
	},
}
