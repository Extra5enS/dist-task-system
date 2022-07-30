package taskBuilder

type TaskType string

const (
	ExtTaskType = "ExtTaskType" // this task should be with with exect
	IntTaskType = "IntTaskType" // this task use golang func
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
	Name     string
	Type     TaskType
	ArgCount int
	Keys     []string
}

var TaskTable = map[string]TaskInfo{
	"hello": {"hello", IntTaskType, 0, []string{}},
	"ls":    {"ls", ExtTaskType, 0, []string{}},
}
