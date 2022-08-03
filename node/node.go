package node

import (
	"log"

	"github.com/Extra5enS/dist-task-system/node/inputer"
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
	"github.com/Extra5enS/dist-task-system/node/utilities"
)

func TermNode() {
	it := inputer.NewInputerTerm()
	c := make(chan taskBuilder.TaskOut)
	end := make(chan interface{})
	counter := utilities.NewCounter(1)
	it.Start(c, end)
	for {
		select {
		case taskOut := <-c:
			if taskOut.E == nil {
				out, e := taskBuilder.TaskExec(taskOut.T)
				taskOut.ReturnAns(out, e)
			} else {
				log.Print(taskOut.E)
			}
		case <-end:
			if counter.IsFinish() {
				return
			}
		}
	}
}
