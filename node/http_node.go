package node

import (
	"fmt"
	"os"

	"github.com/Extra5enS/dist-task-system/node/inputer"
	"github.com/Extra5enS/dist-task-system/node/outputer"
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
	"github.com/Extra5enS/dist-task-system/node/utilities"
	"gopkg.in/yaml.v2"
)

type httpNode struct {
	conf utilities.ServerConfig
	li   []inputer.Inputer
	o    outputer.Outputer
	tw   taskBuilder.TaskWorker

	in  chan taskBuilder.TaskOut
	end chan interface{}
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

	hn.in = make(chan taskBuilder.TaskOut)
	hn.end = make(chan interface{})

	// Create inputer
	it, err := inputer.NewInputerHttp(hn.conf)
	if err != nil {
		fmt.Println(err)
		return httpNode{}, err
	}
	hn.li = append(hn.li, it)

	// Create outputer
	hn.o, err = outputer.NewOutputerHttp(hn.conf)
	if err != nil {
		fmt.Println(err)
		return httpNode{}, err
	}

	// Create last
	hn.tw = taskBuilder.NewTaskWorker(hn.in, hn.end, 1, hn.o)

	return hn, nil
}

func (n httpNode) Start() {
	n.li[0].Start(n.in, n.end)
	n.tw.Start()
}
