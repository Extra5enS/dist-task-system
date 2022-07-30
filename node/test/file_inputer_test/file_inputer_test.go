package inputer

import (
	"testing"

	"github.com/Extra5enS/dist-task-system/node/inputer"
)

func TestFileInputer(t *testing.T) {
	it := inputer.NewInputerFile("test.txt")
	c, end, _ := it.Start()
	for {
		select {
		case task := <-c:
			t.Log(task)
		case <-end:
			return
		}
	}
}
