package inputer

import (
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

type Inputer interface {
	Start() (chan taskBuilder.Task, chan interface{}, error)
	ReturnAns(ans string, e error)
}
