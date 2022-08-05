package outputer

type Outputer interface {
	Get(name string, args []string, incomeAddr string) chan string
	GetByIp(name string, args []string, incomeAddr, addr string) chan string

	AnsCount() int
	OwnAddr() string
}
