package outputer

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gopkg.in/yaml.v2"
)

type outputerHttp struct {
	outAddrs []string
	out      chan string
}

func NewOutputerHttp(conf io.Reader) (outputerHttp, error) {

	mapConf := map[string]interface{}{}
	decoder := yaml.NewDecoder(conf)
	if err := decoder.Decode(mapConf); err != nil {
		return outputerHttp{}, err
	}

	addrs := make([]string, 0)

	if addrs_int, ok := mapConf["addrs"]; ok {
		int_array, ok := addrs_int.([]interface{})
		if !ok {
			return outputerHttp{}, fmt.Errorf("no address config")
		} else {
			for i, addr_int := range int_array {
				addr, ok := addr_int.(string)
				if !ok {
					return outputerHttp{}, fmt.Errorf("wrong arrdes #%d", i)
				}
				addrs = append(addrs, addr)
			}
		}
	}

	return outputerHttp{
		outAddrs: addrs,
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
