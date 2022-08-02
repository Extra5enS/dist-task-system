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
)

type inputerHttp struct {
	taskIdCount taskBuilder.TaskId
	ret         chan string
}

func NewInputerHttp() inputerHttp {
	return inputerHttp{
		taskIdCount: 0,
		ret:         make(chan string),
	}
}

const keyServerAddr = "serverAddr"

func (it inputerHttp) subStart(c chan taskBuilder.TaskOut, end chan interface{}) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		log.Printf("%v\n", ctx)
		command := r.URL.Query().Get("command")
		args := r.URL.Query().Get("args")
		// no command in request
		if command == "" {
			command = "hello"
			args = ""
		}

		newTask := taskBuilder.Task{
			TaskName: command,
			TaskId:   it.taskIdCount,
			Args:     strings.Fields(args),
		}
		it.taskIdCount++

		log.Printf("Task: %v", newTask)
		c <- taskBuilder.TaskOut{T: newTask, E: nil, ret: it.ret}

		out := <-it.ret

		if taskBuilder.TaskTable[command].Type == taskBuilder.IntTaskType {
			out = out + "\n"
		}

		//log.Printf("Res: %v", out)
		io.WriteString(w, out)
	})
	_, cancelCtx := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:        ":4444",
		Handler:     mux,
		BaseContext: nil,
		/*func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},*/
	}
	err := server.ListenAndServe()
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
