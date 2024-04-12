package filters

type vstack struct {
	filter
	inputs   string `args:"inputs"`
	shortest string `args:"shortest"`
}

func VStack() *vstack {
	c := &vstack{}
	c.meta = c
	return c
}

func (me *vstack) Inputs(inputs string) *vstack {
	me.inputs = inputs
	return me
}

func (me *vstack) Shortest(shortest string) *vstack {
	me.shortest = shortest
	return me
}

func (me *vstack) Name() string {
	return "vstack"
}
