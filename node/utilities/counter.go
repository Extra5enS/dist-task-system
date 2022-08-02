package utilities

type counter struct {
	limit int
}

func NewCounter(limit int) counter {
	return counter{
		limit: limit,
	}
}

func (c counter) IsFinish() bool {
	c.limit--
	return c.limit <= 0
}
