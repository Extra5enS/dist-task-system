package main

import (
	"fmt"

	"github.com/Extra5enS/dist-task-system/node"
)

func main() {
	//node.TermNode()
	go func() {
		n, err := node.NewHttpNode("./config/conf_sub_1.yaml")
		if err != nil {
			fmt.Println(err)
			return
		}
		n.Start()
	}()
	n, err := node.NewHttpNode("./config/conf_main.yaml")
	n.Start()
	if err != nil {
		fmt.Println(err)
		return
	}
}
