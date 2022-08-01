package inputer

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
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
		//ctx := r.Context()

		command := r.URL.Query().Get("command")
		args := r.URL.Query().Get("second")

		newTask := taskBuilder.Task{
			TaskName: command,
			TaskId:   it.taskIdCount,
			Args:     strings.Fields(args),
		}
		c <- taskBuilder.TaskOut{T: newTask, E: nil}

		out := <-it.ret
		io.WriteString(w, out)
	})
	ctx, cancelCtx := context.WithCancel(context.Background())
	server := &http.Server{
		Addr:    ":4444",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
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

func (it inputerHttp) Start() (chan taskBuilder.TaskOut, chan interface{}, error) {
	c := make(chan taskBuilder.TaskOut)
	end := make(chan interface{})
	go it.subStart(c, end)
	return c, end, nil
}

func (it inputerHttp) ReturnAns(ans string, e error) {

}