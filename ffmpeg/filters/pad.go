package filters

//填充视频
type pad struct {
	filter
	width  string `args:"w"`
	height string `args:"h"`
	x      string `args:"x"`
	y      string `args:"y"`
	color  string `args:"color"`
	eval   string `args:"eval"`
	aspect string `args:"aspect"`
}

func Pad() *pad {
	c := &pad{}
	c.meta = c
	return c
}
func (me *pad) Width(w string) *pad {
	me.width = w
	return me
}

func (me *pad) Height(h string) *pad {
	me.height = h
	return me
}

func (me *pad) X(x string) *pad {
	me.x = x
	return me
}

func (me *pad) Y(y string) *pad {
	me.y = y
	return me
}

func (me *pad) Color(c string) *pad {
	me.color = c
	return me
}

func (me *pad) Eval(eval string) *pad {
	me.eval = eval
	return me
}

func (me *pad) Aspect(aspect string) *pad {
	me.aspect = aspect
	return me
}

func (me *pad) Name() string {
	return "pad"
}
