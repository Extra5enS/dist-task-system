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
	outAddrs       []string
	outAddrsActive []string
	ownAddrs       string
	out            chan string

	errorAddrs chan string
}

func NewOutputerHttp(conf utilities.ServerConfig) (outputerHttp, error) {
	return outputerHttp{
		outAddrs:       conf.ExtAddrs,
		outAddrsActive: append(conf.ExtAddrs, conf.MyAddr),
		ownAddrs:       conf.MyAddr,
		out:            make(chan string),

		errorAddrs: make(chan string),
	}, nil
}

func (oh outputerHttp) AnsCount() int {
	return len(oh.outAddrs)
}

func (oh outputerHttp) Get(name string, args []string, incomeAddr string) chan string {
	subOut := make(chan string)
	for _, addr := range oh.outAddrsActive {
		if addr == incomeAddr {
			continue
		}
		go func(addr string) {
			requestURL := fmt.Sprintf("http://%s/?name=%s&args=%s&incomeAddr=%s", addr, name, strings.Join(args, "+"), incomeAddr)
			req, err := http.NewRequest(http.MethodGet, requestURL, nil)
			if err != nil {
				log.Printf("%v\n", err)
				oh.errorAddrs <- addr
				subOut <- `{"` + addr + `":"error"}`
				return
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("%v\n", err)
				subOut <- `{"` + addr + `":"error"}`
				return
			}
			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Printf("%v\n", err)
				subOut <- `{"` + addr + `":"error"}`
				return
			}
			subOut <- `{"` + addr + `":"` + string(resBody) + `"}`
		}(addr)
	}
	go func() {
		mergeS := ""
		for _, addr := range oh.outAddrsActive {
			if addr == incomeAddr {
				continue
			}
			mergeS += <-subOut
		}
		oh.out <- mergeS
	}()
	return oh.out
}

func (oh outputerHttp) CleanAddrs() {
}
