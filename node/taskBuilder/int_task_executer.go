package taskBuilder

import "fmt"

type IntTaskExecutor struct {
	// Add something here
}

func (ite IntTaskExecutor) Exec(name string, args []string) (string, error) {
	if exfun, ok := IntTaskExecution[name]; !ok {
		return "", fmt.Errorf("Unknowen command name")
	} else {
		return exfun(args)
	}
}

var IntTaskExecution = map[string](func(args []string) (string, error)){
	"hello": func(args []string) (string, error) {
		fmt.Printf("hello, master!\n")
		return "Done", nil
	},
}
