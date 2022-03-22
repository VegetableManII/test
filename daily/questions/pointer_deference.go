package questions

type items struct {
	Item1 int
	Item2 int
}

type B struct {
	*items
}

func (i *items) UseItem1() int {
	return i.Item1
}
