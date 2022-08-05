package inputer

import (
	"fmt"
	"time"

	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

type inputerRepite struct {
	gen taskBuilder.TaskGenerator
	ret chan string

	command []string
	t       time.Duration
}

func NewInputeRepite(command []string, t time.Duration) (inputerRepite, error) {
	return inputerRepite{
		gen:     taskBuilder.NewTaskGenerator(),
		ret:     make(chan string),
		command: command,
		t:       t,
	}, nil
}

func (it inputerRepite) subStart(c chan taskBuilder.TaskOut, end chan interface{}) {
	for _ = range time.Tick(it.t) {
		if _, ok := taskBuilder.TaskTable[it.command[0]]; !ok {
			fmt.Printf("Unknowen task %s\n", it.command[0])
			continue
		}
		newTask := it.gen.NewTask(it.command[0], it.command[1:], "")
		ret := make(chan string)
		c <- taskBuilder.NewTaskOut(newTask, nil, ret)
		// wait answer
		out := <-ret
		// print result
		fmt.Print(out)
		if taskBuilder.TaskTable[it.command[0]].Type != taskBuilder.SysTaskType {
			fmt.Println()
		}
	}
	// End the thread
	fmt.Println()
	end <- 0
}

func (it inputerRepite) Start(c chan taskBuilder.TaskOut, end chan interface{}) error {
	go it.subStart(c, end)
	return nil
}
