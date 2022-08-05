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
	ownAddr        string
	out            chan string

	errorAddrs chan string
}

func (oh outputerHttp) OwnAddr() string {
	return oh.ownAddr
}

func NewOutputerHttp(conf utilities.ServerConfig) (outputerHttp, error) {
	return outputerHttp{
		outAddrs:       conf.ExtAddrs,
		outAddrsActive: conf.ExtAddrs,
		ownAddr:        conf.MyAddr,
		out:            make(chan string),

		errorAddrs: make(chan string),
	}, nil
}

func (oh outputerHttp) AnsCount() int {
	return len(oh.outAddrsActive)
}

func (oh outputerHttp) Get(name string, args []string, incomeAddr string) chan string {
	subOut := make(chan string)
	//out := make(chan string)
	//var gw sync.WaitGroup
	count := 0
	for _, addr := range oh.outAddrsActive {
		if addr == incomeAddr {
			continue
		}
		//gw.Add(1)
		go func(addr string) {
			requestURL := fmt.Sprintf("http://%s/?name=%s&args=%s&incomeAddr=%s", addr, name, strings.Join(args, "+"), incomeAddr)
			req, err := http.NewRequest(http.MethodGet, requestURL, nil)
			if err != nil {
				log.Printf("%v\n", err)
				oh.errorAddrs <- addr
				subOut <- `{"` + addr + `":"error"}`
				//gw.Done()
				return
			}
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				log.Printf("%v\n", err)
				subOut <- `{"` + addr + `":"error"}`
				//gw.Done()
				return
			}
			resBody, err := ioutil.ReadAll(res.Body)
			if err != nil {
				log.Printf("%v\n", err)
				subOut <- `{"` + addr + `":"error"}`
				//gw.Done()
				return
			}
			//fmt.Println(`{"` + addr + `":"` + string(resBody) + `"}`)
			subOut <- string(resBody)
			//gw.Done()
		}(addr)
		count++
	}
	if count > 0 {
		go func() {
			mergeS := ""
			//gw.Wait()
			for _, addr := range oh.outAddrsActive {
				if addr == incomeAddr {
					continue
				}
				mergeS += <-subOut
			}
			//fmt.Println(mergeS)
			oh.out <- mergeS
		}()
	}
	return oh.out
}

func (oh outputerHttp) GetByIp(name string, args []string, incomeAddr, addr string) chan string {
	//out := make(chan string)
	go func(addr string) {
		requestURL := fmt.Sprintf("http://%s/?name=%s&args=%s&incomeAddr=%s", addr, name, strings.Join(args, "+"), incomeAddr)
		req, err := http.NewRequest(http.MethodGet, requestURL, nil)

		if err != nil {
			log.Printf("%v\n", err)
			oh.errorAddrs <- addr
			oh.out <- `{"` + addr + `":"error"}`
			return
		}
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			log.Printf("%v\n", err)
			oh.out <- `{"` + addr + `":"error"}`
			return
		}
		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("%v\n", err)
			oh.out <- `{"` + addr + `":"error"}`
			return
		}
		if addr == oh.OwnAddr() {
			oh.out <- `{"` + addr + `":"` + string(resBody) + `"}`
		} else {
			oh.out <- string(resBody)
		}
	}(addr)
	return oh.out
}

/*
func (oh outputerHttp) GetByIp(name string, args []string, addrs []string) chan string {

}*/

func (oh outputerHttp) CleanAddrs() {
}
