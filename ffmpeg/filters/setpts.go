package filters

//慢放或快进
type setpts struct {
	filter
	expr string `args:"expr"`
}

func Setpts() *setpts {
	c := &setpts{}
	c.meta = c
	return c
}

//设置速率,如:  2*PTS
func (me *setpts) Expr(expr string) *setpts {
	me.expr = expr
	return me
}

func (me *setpts) Name() string {
	return "setpts"
}
