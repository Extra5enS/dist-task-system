package taskBuilder

type TaskType string

const (
	ExtTaskType = "ExtTaskType" // this task should be with with exect
	IntTaskType = "IntTaskType" // this task use golang func

	Parallel    = "Parallel"
	NonParallel = "NonParallel"
)

// numder of task at a week
type TaskId uint64

type TaskOut struct {
	T Task
	E error
}

type Task struct {
	TaskName string
	TaskId   TaskId
	Args     []string
}

type TaskInfo struct {
	Name string
	Type TaskType
	Keys []string
}

var TaskTable = map[string]TaskInfo{
	"hello": {"hello", IntTaskType, []string{}},
	"ls":    {"ls", ExtTaskType, []string{}},
}
