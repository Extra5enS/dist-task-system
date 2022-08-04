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
	"github.com/Extra5enS/dist-task-system/node/utilities"
)

type inputerHttp struct {
	gen    taskBuilder.TaskGenerator
	server *http.Server
}

func NewInputerHttp(conf utilities.ServerConfig) (inputerHttp, error) {
	/*
		conf - yaml byte array that contain info about Http server
	*/

	if conf.MyAddr == "" {
		return inputerHttp{}, fmt.Errorf("no addr")
	}

	return inputerHttp{
		gen: taskBuilder.NewTaskGenerator(),
		server: &http.Server{
			Addr:        conf.MyAddr,
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
		incomeAddr := r.URL.Query().Get("incomeAddr")

		if name == "" {
			// no command in request
			io.WriteString(w, "No command")
		} else if _, ok := taskBuilder.TaskTable[name]; !ok {
			io.WriteString(w, "Unknowen command: "+name+" "+args)
		} else {
			newTask := it.gen.NewTask(name, strings.Fields(args), incomeAddr)

			log.Printf("Task: %v", newTask)
			ret := make(chan string)

			c <- taskBuilder.NewTaskOut(newTask, nil, ret)
			//go func() {
			out := <-ret
			io.WriteString(w, fmt.Sprintf(`%s`, out))
			//}()

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
