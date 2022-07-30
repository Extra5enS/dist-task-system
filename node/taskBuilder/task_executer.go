package taskBuilder

type TaskExecutor interface {
	Exec(name string, args []string) (string, error)
}
