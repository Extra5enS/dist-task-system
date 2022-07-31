package taskBuilder

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
)

type SysTaskExecutor struct {
	// Add something here
}

func (ste SysTaskExecutor) Exec(name string, args []string) (string, error) {
	os := runtime.GOOS
	switch os {
	case "windows":
		return execWin(name, args)
	case "darwin":
		return execMac(name, args)
	case "linux":
		return execLin(name, args)
	default:
		return "", fmt.Errorf("Not implemented")
	}
}

func execWin(name string, args []string) (string, error) {
	sumCommand := make([]string, 0)
	sumCommand = append(sumCommand, name)
	sumCommand = append(sumCommand, args...)

	commandArg := make([]string, 0)
	commandArg = append(commandArg, "-c")
	commandArg = append(commandArg, strings.Join(sumCommand, " "))
	cmd := exec.Command("bash", commandArg...)

	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	return out.String(), err
}

func execLin(name string, args []string) (string, error) {
	return "", fmt.Errorf("Not implemented")
}

func execMac(name string, args []string) (string, error) {
	return "", fmt.Errorf("Not implemented")
}
