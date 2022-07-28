package taskTable

type TaskId uint64

type Task struct {
	task TaskId
	args []string
}
