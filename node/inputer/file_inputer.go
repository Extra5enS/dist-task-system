package inputer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

type inputerFile struct {
	taskIdCount taskBuilder.TaskId
	fileName    string
	ret         chan string
}

func NewInputerFile(fileName string) inputerFile {
	return inputerFile{
		taskIdCount: 0,
		fileName:    fileName,
		ret:         make(chan string),
	}
}

func (it inputerFile) subStart(c chan taskBuilder.TaskOut, end chan interface{}) {
	file, err := os.Open(it.fileName)
	if err != nil {
		end <- 1
		return
	}
	defer file.Close()

	res_file, err := os.Create("res_" + it.fileName)
	if err != nil {
		end <- 1
		return
	}
	defer res_file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// scan
		command := strings.Split(scanner.Text(), " ")
		if _, ok := taskBuilder.TaskTable[command[0]]; !ok {
			fmt.Fprintf(res_file, "Unknowen task %s\n", command[0])
			continue
		}
		it.taskIdCount++
		newTask := taskBuilder.Task{
			TaskName: command[0],
			TaskId:   it.taskIdCount,
			Args:     command[1:],
		}
		c <- taskBuilder.TaskOut{T: newTask, E: nil, ret: it.ret}

		out := <-it.ret
		fmt.Fprint(res_file, out)
	}
	// End the thread
	end <- 0
}

func (it inputerFile) Start(c chan taskBuilder.TaskOut, end chan interface{}) error {
	go it.subStart(c, end)
	return nil
}
