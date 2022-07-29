package inputer

import (
	"../taskBuilder"
)

type Inputer interface {
	Start() (chan taskBuilder.Task, chan interface{}, error)
}
