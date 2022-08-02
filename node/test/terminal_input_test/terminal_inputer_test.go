package test

import (
	"fmt"
	"testing"

	"github.com/Extra5enS/dist-task-system/node/inputer"
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

func TestTerminalInputer(t *testing.T) {
	it := inputer.NewInputerTerm()
	c := make(chan taskBuilder.TaskOut)
	end := make(chan interface{})
	it.Start(c, end)
	t.Logf("Write term")
	for {
		select {
		case task := <-c:
			fmt.Print(task)
		case <-end:
			break
		}
	}
}
