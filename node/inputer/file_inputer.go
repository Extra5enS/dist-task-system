package inputer

import (
	"bufio"
	"os"
	"strings"

	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

type inputerFile struct {
	taskIdCount taskBuilder.TaskId
	fileName    string
}

func NewInputerFile(fileName string) inputerFile {
	return inputerFile{0, fileName}
}

func (it inputerFile) subStart(c chan taskBuilder.TaskOut, end chan interface{}) {
	file, err := os.Open(it.fileName)
	if err != nil {
		end <- 1
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

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
		c <- taskBuilder.TaskOut{T: newTask, E: nil}
	}
	// End the thread
	end <- 0
}

func (it inputerFile) Start() (chan taskBuilder.TaskOut, chan interface{}, error) {
	c := make(chan taskBuilder.TaskOut)
	end := make(chan interface{})
	go it.subStart(c, end)
	return c, end, nil
}
