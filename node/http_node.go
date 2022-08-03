package node

import (
	"fmt"
	"log"
	"os"

	"github.com/Extra5enS/dist-task-system/node/inputer"
	"github.com/Extra5enS/dist-task-system/node/outputer"
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
	"github.com/Extra5enS/dist-task-system/node/utilities"
	"gopkg.in/yaml.v2"
)

type httpNode struct {
	conf    utilities.ServerConfig
	c       chan taskBuilder.TaskOut
	end     chan interface{}
	counter utilities.Counter
}

func NewHttpNode(conf_name string) (httpNode, error) {
	hn := httpNode{}
	f, err := os.Open(conf_name)
	if err != nil {
		fmt.Println(err)
		return httpNode{}, err
	}
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&hn.conf); err != nil {
		return httpNode{}, err
	}
	f.Close()

	hn.c = make(chan taskBuilder.TaskOut)
	hn.end = make(chan interface{})
	hn.counter = utilities.NewCounter(1)
	return hn, nil
}

func (n httpNode) Start() {
	it, err := inputer.NewInputerHttp(n.conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputer, err := outputer.NewOutputerHttp(n.conf)
	if err != nil {
		fmt.Println(err)
		return
	}

	it.Start(n.c, n.end)
	for {
		select {
		case task := <-n.c:
			if task.E == nil {
				out, e := taskBuilder.TaskExec(task.T, outputer)
				task.ReturnAns(out, e)
			} else {
				log.Print(task.E)
			}
		case <-n.end:
			if n.counter.IsFinish() {
				return
			}
		}
	}
}
