package node

import (
	"fmt"
	"os"
	"time"

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
	out chan taskBuilder.Ans
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
	hn.out = make(chan taskBuilder.Ans)
	hn.end = make(chan interface{})

	// Create inputer
	it, err := inputer.NewInputerHttp(hn.conf)
	if err != nil {
		fmt.Println(err)
		return httpNode{}, err
	}
	hn.li = append(hn.li, it)

	command := []string{"foreach", "nothing"}
	ir, err := inputer.NewInputeRepite(
		command,
		10*time.Second,
	)
	hn.li = append(hn.li, ir)

	// Create outputer
	hn.o, err = outputer.NewOutputerHttp(hn.conf)
	if err != nil {
		fmt.Println(err)
		return httpNode{}, err
	}

	// Create last
	hn.tw = taskBuilder.NewTaskWorker(hn.in, hn.out, hn.end, len(hn.li), hn.o)

	return hn, nil
}

func (n httpNode) Start() {
	for _, l := range n.li {
		l.Start(n.in, n.end)
	}
	n.tw.Start()
}
