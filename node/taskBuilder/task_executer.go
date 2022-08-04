package taskBuilder

import (
	"log"

	"github.com/Extra5enS/dist-task-system/node/outputer"
	"github.com/Extra5enS/dist-task-system/node/utilities"
)

type TaskExecutor interface {
	Exec(name string, args []string) (string, error)
}

func TaskExec(t Task, o outputer.Outputer) (string, error) {
	var te TaskExecutor
	switch TaskTable[t.Name].Type {
	case IntTaskType:
		te = IntTaskExecutor{}
	case ExtTaskType:
		te = NewExtTaskExecutor(o)
	case SysTaskType:
		te = SysTaskExecutor{}
	}
	out, e := te.Exec(t.Name, t.Args)
	if e != nil {
		return out, e
	} else {
		return out, e
	}
}

//
type TaskWorker struct {
	in      chan TaskOut
	o       outputer.Outputer
	end     chan interface{}
	counter utilities.Counter
}

func NewTaskWorker(in chan TaskOut, end chan interface{}, limit int, o outputer.Outputer) TaskWorker {
	return TaskWorker{
		in:      in,
		o:       o,
		end:     end,
		counter: utilities.NewCounter(limit),
	}
}

func (tw TaskWorker) Start() {
	for {
		select {
		case task := <-tw.in:
			if task.E == nil {
				out, e := TaskExec(task.T, tw.o)
				task.ReturnAns(out, e)
			} else {
				log.Print(task.E)
			}
		case <-tw.end:
			if tw.counter.IsFinish() {
				return
			}
		}
	}
}
