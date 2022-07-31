package inputer

import "github.com/Extra5enS/dist-task-system/node/taskBuilder"

type inputerHppt struct {
	taskIdCount taskBuilder.TaskId
	ret         chan string
}

func NewInputerHppt() inputerHppt {
	return inputerHppt{
		taskIdCount: 0,
		ret:         make(chan string),
	}
}

func (it inputerHppt) subStart(c chan taskBuilder.TaskOut, end chan interface{}) {

}

func (it inputerHppt) Start() (chan taskBuilder.TaskOut, chan interface{}, error) {
	c := make(chan taskBuilder.TaskOut)
	end := make(chan interface{})
	go it.subStart(c, end)
	return c, end, nil
}

func (it inputerHppt) ReturnAns(ans string, e error) {

}
