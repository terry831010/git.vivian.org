package filters

//旋转
type rotate struct {
	filter
	angle     string `args:"a"`
	out_w     string `args:"ow"`
	out_h     string `args:"oh"`
	fillcolor string `args:"c"`
	bilinear  string `args:"bilinear"`
	enable    string `args:"enable"`
}

func Rotate() *rotate {
	c := &rotate{}
	c.meta = c
	return c
}

func (me *rotate) Angle(angle string) *rotate {
	me.angle = angle
	return me
}

func (me *rotate) OutW(ow string) *rotate {
	me.out_w = ow
	return me
}

func (me *rotate) OutH(oh string) *rotate {
	me.out_h = oh
	return me
}

func (me *rotate) FillColor(c string) *rotate {
	me.fillcolor = c
	return me
}

func (me *rotate) Bilinear(bilinear string) *rotate {
	me.bilinear = bilinear
	return me
}

func (me *rotate) Timeline(enable string) *rotate {
	me.enable = enable
	return me
}

func (me *rotate) Name() string {
	return "rotate"
}
