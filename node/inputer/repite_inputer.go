package inputer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

type inputerRepite struct {
	gen taskBuilder.TaskGenerator
	ret chan string
}

func NewInputeRepite() inputerTerm {
	return inputerTerm{
		gen: taskBuilder.NewTaskGenerator(),
		ret: make(chan string),
	}
}

func (it inputerRepite) subStart(c chan taskBuilder.TaskOut, end chan interface{}) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("user> ")
	for scanner.Scan() {
		// scan
		command := strings.Split(scanner.Text(), " ")

		if _, ok := taskBuilder.TaskTable[command[0]]; !ok {
			fmt.Printf("Unknowen task %s\n", command[0])
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
	}
	// End the thread
	fmt.Println()
	end <- 0
}

func (it inputerRepite) Start(c chan taskBuilder.TaskOut, end chan interface{}) error {
	go it.subStart(c, end)
	return nil
}
