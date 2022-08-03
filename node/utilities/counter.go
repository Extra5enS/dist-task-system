package utilities

type Counter struct {
	limit int
}

func NewCounter(limit int) Counter {
	return Counter{
		limit: limit,
	}
}

func (c Counter) IsFinish() bool {
	c.limit--
	return c.limit <= 0
}
