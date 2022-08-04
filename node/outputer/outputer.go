package outputer

type Outputer interface {
	Get(name string, args []string, incomeAddr string) chan string
	AnsCount() int
}
