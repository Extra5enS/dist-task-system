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
		newTask := taskBuilder.Task{
			TaskName: command[0],
			TaskId:   it.taskIdCount,
			Args:     command[1:],
		}
		it.taskIdCount++

		c <- taskBuilder.TaskOut{T: newTask, E: nil, ret: it.ret}
		// wait answer
		out := <-it.ret
		// print result
		fmt.Print(out)
		if taskBuilder.TaskTable[command[0]].Type != taskBuilder.SysTaskType {
			fmt.Println()
		}
		fmt.Print("user> ")
	}
	// End the thread
	fmt.Println()
	end <- 0
}

func (it inputerTerm) Start(c chan taskBuilder.TaskOut, end chan interface{}) error {
	go it.subStart(c, end)
	return nil
}
