package outputer

type Outputer interface {
	Get(name string, args []string) chan string
	AnsCount() int
}
