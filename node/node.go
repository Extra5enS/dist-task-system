package main

import (
	"fmt"

	"github.com/Extra5enS/dist-task-system/node/inputer"
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

func taskExec(t taskBuilder.Task) error {
	var te taskBuilder.TaskExecutor
	if taskBuilder.TaskTable[t.TaskName].Type == taskBuilder.IntTaskType {
		te = taskBuilder.IntTaskExecutor{}
	} else {
		return fmt.Errorf("ExtTaskExecutor didn't impl")
	}
	_, e := te.Exec(t.TaskName, t.Args)
	if e != nil {
		return e
	} else {
		return e
	}
}

func main() {
	it := inputer.NewInputerTerm()
	c, end, _ := it.Start()
	for {
		select {
		case task := <-c:
			if task.E == nil {
				taskExec(task.T)
			} else {
				fmt.Print(task.E, "\n")
			}
		case <-end:
			return
		}
	}
}
