package taskBuilder

import (
	"bytes"
	"os/exec"
)

type SysTaskExecutor struct {
	// Add something here
}

func (ste SysTaskExecutor) Exec(name string, args []string) (string, error) {
	commandArg := make([]string, 0)
	commandArg = append(commandArg, "-c")
	commandArg = append(commandArg, name)
	commandArg = append(commandArg, args...)
	cmd := exec.Command("bash", commandArg...)
	//fmt.Print(name, "\n")
	//cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}
