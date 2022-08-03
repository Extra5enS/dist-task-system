package node

import (
	"log"
	"os"

	"github.com/Extra5enS/dist-task-system/node/inputer"
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
	"github.com/Extra5enS/dist-task-system/node/utilities"
)

func HttpNode(conf_name string) {
	f, err := os.Open(conf_name)
	if err != nil {
		return
	}
	it, _ := inputer.NewInputerHttp(f)
	c := make(chan taskBuilder.TaskOut)
	end := make(chan interface{})
	counter := utilities.NewCounter(1)
	it.Start(c, end)
	for {
		select {
		case task := <-c:
			if task.E == nil {
				out, e := taskBuilder.TaskExec(task.T)
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
