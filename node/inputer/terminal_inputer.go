package inputer

import (
	"bufio"
	"os"
	"strings"

	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

type InputerTerm struct {
	taskIdCount taskBuilder.TaskId
}

func (it InputerTerm) subStart(c chan taskBuilder.Task, end chan interface{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// scan
		command := strings.Split(scanner.Text(), " ")
		if _, ok := taskBuilder.TaskTable[command[0]]; !ok {
			continue
		}
		it.taskIdCount++
		newTask := taskBuilder.Task{
			TaskName: command[0],
			TaskId:   it.taskIdCount,
			Keys:     []string{},
			Args:     command[1:],
		}
		c <- newTask
	}
	// End the thread
	end <- 0
}

func (it InputerTerm) Start() (chan taskBuilder.Task, chan interface{}, error) {
	c := make(chan taskBuilder.Task)
	end := make(chan interface{})
	go it.subStart(c, end)
	return c, end, nil
}
