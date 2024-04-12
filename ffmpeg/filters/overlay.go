package filters

//画中画
type overlay struct {
	filter
	x          string `args:"x"`
	y          string `args:"y"`
	eof_action string `args:"eof_action"`
	eval       string `args:"eval"`
	shortest   string `args:"shortest"`
	format     string `args:"format"`
	repeatlast string `args:"repeatlast"`
	alpha      string `args:"alpha"`
	enable     string `args:"enable"`
}

func Overlay() *overlay {
	c := &overlay{}
	c.meta = c
	return c
}

func (me *overlay) X(x string) *overlay {
	me.x = x
	return me
}

func (me *overlay) Y(y string) *overlay {
	me.y = y
	return me
}

func (me *overlay) EofAction(eof_action string) *overlay {
	me.eof_action = eof_action
	return me
}
func (me *overlay) Eval(eval string) *overlay {
	me.eval = eval
	return me
}
func (me *overlay) Shortest(shortest string) *overlay {
	me.shortest = shortest
	return me
}
func (me *overlay) Format(format string) *overlay {
	me.format = format
	return me
}
func (me *overlay) Repeatlast(repeatlast string) *overlay {
	me.repeatlast = repeatlast
	return me
}
func (me *overlay) Alpha(alpha string) *overlay {
	me.alpha = alpha
	return me
}

func (me *overlay) Timeline(enable string) *overlay {
	me.enable = enable
	return me
}

func (me *overlay) Name() string {
	return "overlay"
}
