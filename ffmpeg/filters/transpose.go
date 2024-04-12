package filters

//整体转置
type transpose struct {
	filter
	dir         string `args:"dir"`
	passthrough string `args:"passthrough"`
}

func Transpose() *transpose {
	c := &transpose{}
	c.meta = c
	return c
}

func (me *transpose) Dir(dir string) *transpose {
	me.dir = dir
	return me
}

func (me *transpose) Passthrough(passthrough string) *transpose {
	me.passthrough = passthrough
	return me
}

func (me *transpose) Name() string {
	return "transpose"
}
