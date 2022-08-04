package inputer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

type inputerTerm struct {
	gen taskBuilder.TaskGenerator
	ret chan string
}

func NewInputerTerm() inputerTerm {
	return inputerTerm{
		gen: taskBuilder.NewTaskGenerator(),
		ret: make(chan string),
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
		newTask := it.gen.NewTask(command[0], command[1:], "")
		c <- taskBuilder.NewTaskOut(newTask, nil, it.ret)
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
