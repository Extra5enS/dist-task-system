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

type TaskGenerator struct {
	counter TaskId
}

func NewTaskGenerator() TaskGenerator {
	return TaskGenerator{counter: 0}
}

type Task struct {
	Name string
	Args []string

	TaskId     TaskId
	IncomeAddr string
}

func (tg TaskGenerator) NewTask(name string, args []string, addr string) Task {
	t := Task{
		Name: name,
		Args: args,

		TaskId:     tg.counter,
		IncomeAddr: addr,
	}
	tg.counter++
	return t
}

type TaskOut struct {
	T Task
	E error

	Ret chan string
}

func NewTaskOut(t Task, e error, ret chan string) TaskOut {
	return TaskOut{T: t, E: e, Ret: ret}
}

func (t TaskOut) ReturnAns(ans string, e error) {
	if e != nil {
		t.Ret <- ans
	} else {
		t.Ret <- ans
	}
}

type TaskInfo struct {
	Name string
	Type TaskType
	Keys []string
}

var TaskTable = map[string]TaskInfo{
	"hello":   {"hello", IntTaskType, []string{}},
	"sum":     {"sum", IntTaskType, []string{}},
	"ls":      {"ls", SysTaskType, []string{}},
	"dir":     {"dir", SysTaskType, []string{}},
	"rm":      {"rm", SysTaskType, []string{}},
	"mkdir":   {"mkdir", SysTaskType, []string{}},
	"cat":     {"cat", SysTaskType, []string{}},
	"pwd":     {"pwd", SysTaskType, []string{}},
	"echo":    {"echo", SysTaskType, []string{}},
	"export":  {"export", SysTaskType, []string{}},
	"env":     {"env", SysTaskType, []string{}},
	"foreach": {"foreach", ExtTaskType, []string{}},
}
