package test

import (
	"fmt"
	"testing"

	"github.com/Extra5enS/dist-task-system/node/inputer"
)

func TestTerminalInputer(t *testing.T) {
	it := inputer.NewInputerTerm()
	c, end, _ := it.Start()
	t.Logf("Write term")
	for {
		select {
		case task := <-c:
			fmt.Print(task)
		case <-end:
			break
		default:
			t.Errorf("Unclear error")
			break
		}
	}
}
