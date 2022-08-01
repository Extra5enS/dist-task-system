package node

import (
	"log"

	"github.com/Extra5enS/dist-task-system/node/inputer"
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

func HttpNode() {
	it := inputer.NewInputerHttp()
	c, end, _ := it.Start()
	for {
		select {
		case task := <-c:
			if task.E == nil {
				out, e := taskBuilder.TaskExec(task.T)
				it.ReturnAns(out, e)
			} else {
				log.Print(task.E)
			}
		case <-end:
			return
		}
	}
}
