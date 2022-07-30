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
			if task.E == nil {
				t.Logf("%v\n", task.T)
			} else {
				t.Logf("Input Error %v\n", task.E)
			}
		case <-end:
			return
		}
	}
}
