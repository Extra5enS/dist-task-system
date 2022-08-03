package inputer

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
	"github.com/go-yaml/yaml"
)

type inputerHttp struct {
	gen    taskBuilder.TaskGenerator
	ret    chan string
	server *http.Server
}

func NewInputerHttp(conf io.Reader) (inputerHttp, error) {
	/*
		conf - yaml byte array that contain info about Http server
	*/

	mapConf := map[string]string{}
	decoder := yaml.NewDecoder(conf)
	if err := decoder.Decode(mapConf); err != nil {
		return inputerHttp{}, err
	}

	var (
		addr string
		ok   bool
	)

	if addr, ok = mapConf["addr"]; !ok {
		return inputerHttp{}, fmt.Errorf("no address config")
	}

	return inputerHttp{
		gen: taskBuilder.NewTaskGenerator(),
		ret: make(chan string),
		server: &http.Server{
			Addr:        addr,
			Handler:     nil,
			BaseContext: nil,
			/*func(l net.Listener) context.Context {
				ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
				return ctx
			},*/
		},
	}, nil
}

func (it inputerHttp) subStart(c chan taskBuilder.TaskOut, end chan interface{}) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log.Printf("%v\n", ctx)
		name := r.URL.Query().Get("name")
		args := r.URL.Query().Get("args")
		// no command in request
		if name == "" {
			io.WriteString(w, "No command")
		} else if _, ok := taskBuilder.TaskTable[name]; !ok {
			io.WriteString(w, "Unknowen command: "+name+" "+args)
		} else {
			newTask := it.gen.NewTask(name, strings.Fields(args))

			log.Printf("Task: %v", newTask)
			c <- taskBuilder.NewTaskOut(newTask, nil, it.ret)

			out := <-it.ret

			if taskBuilder.TaskTable[name].Type == taskBuilder.IntTaskType {
				out = out + "\n"
			}

			io.WriteString(w, out)
		}
	})

	_, cancelCtx := context.WithCancel(context.Background())
	it.server.Handler = mux

	err := it.server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server one closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server one: %s\n", err)
	}
	cancelCtx()
	end <- 0
}

func (it inputerHttp) Start(c chan taskBuilder.TaskOut, end chan interface{}) error {
	go it.subStart(c, end)
	return nil
}
