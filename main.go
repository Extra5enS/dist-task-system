package main

import "github.com/Extra5enS/dist-task-system/node"

func main() {
	//node.TermNode()
	go func() {
		node.HttpNode("./config/conf_sub_1.yaml")
	}()
	node.HttpNode("./config/conf_main.yaml")

}
