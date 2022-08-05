package taskBuilder

import (
	"log"

	"github.com/Extra5enS/dist-task-system/node/outputer"
	"github.com/Extra5enS/dist-task-system/node/utilities"
)

type TaskExecutor interface {
	Exec(t Task) (string, error)
}

func (tw TaskWorker) TaskExec(t TaskOut) {
	go func() {
		var te TaskExecutor
		switch TaskTable[t.T.Name].Type {
		case IntTaskType:
			te = IntTaskExecutor{}
		case ExtTaskType:
			te = NewExtTaskExecutor(tw.o)
		case SysTaskType:
			te = SysTaskExecutor{}
		}
		outString, e := te.Exec(t.T)
		if e != nil {
			tw.out <- Ans{Out: e.Error(), RetChan: t.Ret, AnsTask: t.T}
		} else {
			tw.out <- Ans{Out: outString, RetChan: t.Ret, AnsTask: t.T}
		}
	}()
}

//
type TaskWorker struct {
	in      chan TaskOut
	out     chan Ans
	o       outputer.Outputer
	end     chan interface{}
	counter utilities.Counter
}

func NewTaskWorker(in chan TaskOut, out chan Ans, end chan interface{}, limit int, o outputer.Outputer) TaskWorker {
	return TaskWorker{
		in:      in,
		out:     out,
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
				//fmt.Println("taskOut:", task)
				tw.TaskExec(task)
			} else {
				log.Print(task.E)
			}
		case asn := <-tw.out:
			//fmt.Println("Ans:", asn)
			asn.RetChan <- asn.Out
		case <-tw.end:
			if tw.counter.IsFinish() {
				return
			}
		}
	}
}
