package taskBuilder

type TaskType string

const (
	ExtTaskType = "ExtTaskType"
	IntTaskType = "IntTaskType" // this task use golang func
	SysTaskType = "SysTaskType" // this task should be with with exect

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
	"ls":    {"ls", SysTaskType, []string{}},
	"dir":   {"dir", SysTaskType, []string{}},
	//"cd":    {"cd", SysTaskType, []string{}},
	"rm":    {"rm", SysTaskType, []string{}},
	"mkdir": {"mkdir", SysTaskType, []string{}},
	"cat":   {"cat", SysTaskType, []string{}},
	"pwd":   {"pwd", SysTaskType, []string{}},
}
