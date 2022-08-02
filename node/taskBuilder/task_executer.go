package taskBuilder

import "fmt"

type TaskExecutor interface {
	Exec(name string, args []string) (string, error)
}

func TaskExec(t Task) (string, error) {
	var te TaskExecutor
	switch TaskTable[t.Name].Type {
	case IntTaskType:
		te = IntTaskExecutor{}
	case ExtTaskType:
		return "", fmt.Errorf("ExtTaskExecutor didn't impl")
	case SysTaskType:
		te = SysTaskExecutor{}
	}
	out, e := te.Exec(t.Name, t.Args)
	if e != nil {
		return out, e
	} else {
		return out, e
	}
}
