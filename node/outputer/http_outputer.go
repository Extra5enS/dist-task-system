package outputer

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Extra5enS/dist-task-system/node/utilities"
)

type outputerHttp struct {
	outAddrs []string
	out      chan string
}

func NewOutputerHttp(conf utilities.ServerConfig) (outputerHttp, error) {
	return outputerHttp{
		outAddrs: conf.ExtAddrs,
		out:      make(chan string),
	}, nil
}

func (oh outputerHttp) AnsCount() int {
	return len(oh.outAddrs)
}

func (oh outputerHttp) Get(name string, args []string) chan string {
	for _, addr := range oh.outAddrs {
		go func(addr string) {
			requestURL := fmt.Sprintf("http://%s/?name=%s&args=%s", addr, name, strings.Join(args, "+"))
			req, err := http.NewRequest(http.MethodGet, requestURL, nil)
			if err != nil {
				log.Printf("%v\n", err)
				oh.out <- "error"
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("%v\n", err)
				oh.out <- "error"
			}
			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Printf("%v\n", err)
				oh.out <- "error"
			}
			oh.out <- string(resBody)
		}(addr)
	}
	return oh.out
}
