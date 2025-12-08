package card

type Card struct {
	Body Body
}

func (c *Card) Lockup(in string) (out string) {
	c.Body.ReadByLine(func(line string) (exit bool) {
		item := (&Item{}).Parser(line)

		if item.Key == in {
			out = item.Value
			return true
		}
		return
	})
	return
}
