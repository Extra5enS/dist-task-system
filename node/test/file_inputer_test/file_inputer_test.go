package inputer

import (
	"testing"

	"github.com/Extra5enS/dist-task-system/node/inputer"
	"github.com/Extra5enS/dist-task-system/node/taskBuilder"
)

func TestFileInputer(t *testing.T) {
	it := inputer.NewInputerFile("test.txt")
	c := make(chan taskBuilder.TaskOut)
	end := make(chan interface{})
	it.Start(c, end)
	for {
		select {
		case task := <-c:
			if task.E == nil {
				t.Logf("%v\n", task.T)
				it.ReturnAns(task.T.TaskName+"\n", nil)
			} else {
				t.Logf("Input Error %v\n", task.E)
			}
		case end_code := <-end:
			if end_code != 0 {
				t.Errorf("File problem")
			}
			return
		}
	}
}
