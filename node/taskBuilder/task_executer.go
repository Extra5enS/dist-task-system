package taskBuilder

import (
	"log"

	"github.com/Extra5enS/dist-task-system/node/outputer"
	"github.com/Extra5enS/dist-task-system/node/utilities"
)

type TaskExecutor interface {
	Exec(t Task) (string, error)
}

func TaskExec(t TaskOut, o outputer.Outputer) {
	go func() {
		var te TaskExecutor
		switch TaskTable[t.T.Name].Type {
		case IntTaskType:
			te = IntTaskExecutor{}
		case ExtTaskType:
			te = NewExtTaskExecutor(o)
		case SysTaskType:
			te = SysTaskExecutor{}
		}
		out, e := te.Exec(t.T)
		t.ReturnAns(out, e)
	}()
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
				TaskExec(task, tw.o)
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
