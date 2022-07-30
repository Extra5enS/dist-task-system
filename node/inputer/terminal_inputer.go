package inputer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

type inputerTerm struct {
	taskIdCount taskBuilder.TaskId
}

func NewInputerTerm() inputerTerm {
	return inputerTerm{0}
}

func (it inputerTerm) subStart(c chan taskBuilder.TaskOut, end chan interface{}) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// scan
		command := strings.Split(scanner.Text(), " ")
		if _, ok := taskBuilder.TaskTable[command[0]]; !ok {
			c <- taskBuilder.TaskOut{E: fmt.Errorf("Unknowen task %s", command[0])}
			continue
		}
		it.taskIdCount++
		newTask := taskBuilder.Task{
			TaskName: command[0],
			TaskId:   it.taskIdCount,
			Args:     command[1:],
		}
		c <- taskBuilder.TaskOut{T: newTask, E: nil}
	}
	// End the thread
	end <- 0
}

func (it inputerTerm) Start() (chan taskBuilder.TaskOut, chan interface{}, error) {
	c := make(chan taskBuilder.TaskOut)
	end := make(chan interface{})
	go it.subStart(c, end)
	return c, end, nil
}
