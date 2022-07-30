package inputer

import (
	"fmt"
	"testing"

	"github.com/Extra5enS/dist-task-system/node/inputer"
)

func TestFileInputer(t *testing.T) {
	it := inputer.NewInputerFile("test.txt")
	c, end, _ := it.Start()
	for {
		select {
		case task := <-c:
			fmt.Print(task)
		case <-end:
			return
		default:
			t.Errorf("Unclear error")
			return
		}
	}
}
