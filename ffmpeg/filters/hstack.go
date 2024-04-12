package filters

//横向排列
type hstack struct {
	filter
	inputs   string `args:"inputs"`
	shortest string `args:"shortest"`
}

func HStack() *hstack {
	c := &hstack{}
	c.meta = c
	return c
}

func (me *hstack) Inputs(inputs string) *hstack {
	me.inputs = inputs
	return me
}

func (me *hstack) Shortest(shortest string) *hstack {
	me.shortest = shortest
	return me
}

func (me *hstack) Name() string {
	return "hstack"
}
