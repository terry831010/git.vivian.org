package filters

// 产生或逆转渐晕效应
type vignette struct {
	filter
	angle    string `args:"a"`
	x0       string `args:"x0"`
	y0       string `args:"y0"`
	mode     string `args:"mode"`
	eval     string `args:"eval"`
	dither   string `args:"dither"`
	timeline string `args:"enable"`
}

func Vignette() *vignette {
	c := &vignette{}
	c.meta = c
	return c
}

func (me *vignette) Angle(angle string) *vignette {
	me.angle = angle
	return me
}

func (me *vignette) SetX0(x0 string) *vignette {
	me.x0 = x0
	return me
}

func (me *vignette) SetY0(y0 string) *vignette {
	me.y0 = y0
	return me
}

func (me *vignette) Mode(mode string) *vignette {
	me.mode = mode
	return me
}
func (me *vignette) Eval(eval string) *vignette {
	me.eval = eval
	return me
}
func (me *vignette) Dither(dither string) *vignette {
	me.dither = dither
	return me
}

func (me *vignette) Timeline(enable string) *vignette {
	me.timeline = enable
	return me
}
func (me *vignette) Name() string {
	return "vignette"
}
