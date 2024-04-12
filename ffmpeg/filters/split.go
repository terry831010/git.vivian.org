package filters

type split struct {
	filter
	outputs string `args:"outputs"`
}

func Split() *split {
	c := &split{}
	c.meta = c
	return c
}

func (me *split) SetOutPuts(n string) *split {
	me.outputs = n
	return me
}

func (me *split) Name() string {
	return "split"
}
