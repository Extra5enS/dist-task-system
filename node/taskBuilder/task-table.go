package taskBuilder

// numder of task at a week
type TaskId uint64

type Task struct {
	TaskName string
	TaskId   TaskId
	Keys     []string
	Args     []string
}

type TaskInfo struct {
	Name     string
	ArgCount int
	Keys     []string
}

var TaskTable = map[string]TaskInfo{
	"hello": {"hello", 0, []string{}},
	//"lookup": {"lookup", 1, []string{"lRa"}},
	//"read":   {"read", 1, []string{""}},
}
