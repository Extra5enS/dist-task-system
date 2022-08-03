package node

import (
	"fmt"
	"log"
	"os"

	"github.com/Extra5enS/dist-task-system/node/inputer"
	"github.com/Extra5enS/dist-task-system/node/outputer"
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
	"github.com/Extra5enS/dist-task-system/node/utilities"
)

func HttpNode(conf_name string) {
	f, err := os.Open(conf_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	it, err := inputer.NewInputerHttp(f)
	f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	c := make(chan taskBuilder.TaskOut)
	end := make(chan interface{})
	counter := utilities.NewCounter(1)

	f, err = os.Open(conf_name)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputer, err := outputer.NewOutputerHttp(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Close()

	it.Start(c, end)
	for {
		select {
		case task := <-c:
			if task.E == nil {
				out, e := taskBuilder.TaskExec(task.T, outputer)
				task.ReturnAns(out, e)
			} else {
				log.Print(task.E)
			}
		case <-end:
			if counter.IsFinish() {
				return
			}
		}
	}
}
