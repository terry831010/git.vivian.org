package filters

type xfade struct {
	filter
	transition string `args:"transition"`
	duration   string `args:"duration"`
	offset     string `args:"offset"`
	expr       string `args:"expr"`
}

func Xfade() *xfade {
	c := &xfade{}
	c.meta = c
	return c
}

func (me *xfade) Transition(transition string) *xfade {
	me.transition = transition
	return me
}

func (me *xfade) Duration(d string) *xfade {
	me.duration = d
	return me
}

func (me *xfade) Offset(offset string) *xfade {
	me.offset = offset
	return me
}

func (me *xfade) Expr(expr string) *xfade {
	me.expr = expr
	return me
}

func (me *xfade) Name() string {
	return "xfade"
}
