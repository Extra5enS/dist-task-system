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
	ret         chan string
}

func NewInputerTerm() inputerTerm {
	return inputerTerm{
		taskIdCount: 0,
		ret:         make(chan string),
	}
}

func (it inputerTerm) subStart(c chan taskBuilder.TaskOut, end chan interface{}) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("user> ")
	for scanner.Scan() {
		// scan
		command := strings.Split(scanner.Text(), " ")

		if _, ok := taskBuilder.TaskTable[command[0]]; !ok {
			fmt.Printf("Unknowen task %s\n", command[0])
			fmt.Print("user> ")
			continue
		}
		it.taskIdCount++
		newTask := taskBuilder.Task{
			TaskName: command[0],
			TaskId:   it.taskIdCount,
			Args:     command[1:],
		}
		c <- taskBuilder.TaskOut{T: newTask, E: nil}
		// wait answer
		out := <-it.ret
		fmt.Print(out)
		fmt.Print("user> ")
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

func (it inputerTerm) ReturnAns(ans string, e error) {
	if e != nil {
		it.ret <- e.Error()
	} else {
		it.ret <- ans
	}
}
