package main

import (
	"fmt"

	"github.com/Extra5enS/dist-task-system/node/inputer"
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

func taskExec(t taskBuilder.Task) (string, error) {
	var te taskBuilder.TaskExecutor
	switch taskBuilder.TaskTable[t.TaskName].Type {
	case taskBuilder.IntTaskType:
		te = taskBuilder.IntTaskExecutor{}
	case taskBuilder.ExtTaskType:
		return "", fmt.Errorf("ExtTaskExecutor didn't impl")
	case taskBuilder.SysTaskType:
		te = taskBuilder.SysTaskExecutor{}
	}
	out, e := te.Exec(t.TaskName, t.Args)
	if e != nil {
		return out, e
	} else {
		return out, e
	}
}

func main() {
	it := inputer.NewInputerTerm()
	c, end, _ := it.Start()
	for {
		select {
		case task := <-c:
			if task.E == nil {
				out, e := taskExec(task.T)
				it.ReturnAns(out, e)
			} else {
				it.ReturnAns("", task.E)
			}
		case <-end:
			return
		}
	}
}
