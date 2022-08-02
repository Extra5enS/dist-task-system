package inputer

import (
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

type Inputer interface {
	Start(chan taskBuilder.TaskOut, chan interface{}) error
}
