package filters

//连接过滤器，暂未被shuttle封装
type concat struct {
	filter
	n      string `args:"n"`
	v      string `args:"v"`
	a      string `args:"a"`
	unsafe string `args:"unsafe"`
}

func Concat() *concat {
	c := &concat{}
	c.meta = c
	return c
}

func (me *concat) SetN(n string) *concat {
	me.n = n
	return me
}

func (me *concat) SetV(v string) *concat {
	me.v = v
	return me
}

func (me *concat) SetA(a string) *concat {
	me.a = a
	return me
}

func (me *concat) UnSafe(unsafe string) *concat {
	me.unsafe = unsafe
	return me
}

func (me *concat) Name() string {
	return "concat"
}
