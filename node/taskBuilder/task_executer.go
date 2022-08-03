package taskBuilder

import (
	"github.com/Extra5enS/dist-task-system/node/outputer"
)

type TaskExecutor interface {
	Exec(name string, args []string) (string, error)
}

func TaskExec(t Task, o outputer.Outputer) (string, error) {
	var te TaskExecutor
	switch TaskTable[t.Name].Type {
	case IntTaskType:
		te = IntTaskExecutor{}
	case ExtTaskType:
		te = NewExtTaskExecutor(o)
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
