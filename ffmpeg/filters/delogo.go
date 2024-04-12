package filters

//指定区域模糊
type delogo struct {
	filter
	x      string `args:"x"`
	y      string `args:"y"`
	width  string `args:"w"`
	height string `args:"h"`
	show   string `args:"show"`
	enable string `args:"enable"`
}

func DeLogo() *delogo {
	c := &delogo{}
	c.meta = c
	return c
}

func (me *delogo) SetX(x string) *delogo {
	me.x = x
	return me
}

func (me *delogo) SetY(y string) *delogo {
	me.y = y
	return me
}

func (me *delogo) Width(w string) *delogo {
	me.width = w
	return me
}

func (me *delogo) Height(h string) *delogo {
	me.height = h
	return me
}

func (me *delogo) Show(show string) *delogo {
	me.show = show
	return me
}

func (me *delogo) Timeline(enable string) *delogo {
	me.enable = enable
	return me
}

func (me *delogo) Name() string {
	return "delogo"
}
