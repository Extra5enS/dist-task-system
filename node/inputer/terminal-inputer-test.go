package inputer

import (
	"fmt"
	"testing"
)

func TestTerminalInputer(t *testing.T) {
	it := InputerTerm{0}
	c, end, _ := it.Start()
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
