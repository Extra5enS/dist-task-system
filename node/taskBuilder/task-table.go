package taskBuilder

type TaskId uint64

type Task struct {
	task TaskId
	args []string
}
