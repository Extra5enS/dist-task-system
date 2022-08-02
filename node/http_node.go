package node

import (
	"log"

	"github.com/Extra5enS/dist-task-system/node/inputer"
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
	"github.com/Extra5enS/dist-task-system/node/utilities"
)

func HttpNode() {
	it := inputer.NewInputerHttp()
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
