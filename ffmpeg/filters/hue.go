package filters

//色调调整
type hue struct {
	filter
	hd     string `args:"h"`
	s      string `args:"s"`
	hr     string `args:"H"`
	b      string `args:"b"`
	enable string `args:"enable"`
}

func Hue() *hue {
	c := &hue{}
	c.meta = c
	return c
}

func (me *hue) Degrees(hd string) *hue {
	me.hd = hd
	return me
}

func (me *hue) Saturation(s string) *hue {
	me.s = s
	return me
}

func (me *hue) Radians(hr string) *hue {
	me.hr = hr
	return me
}

func (me *hue) Brightness(b string) *hue {
	me.b = b
	return me
}

func (me *hue) Timeline(enable string) *hue {
	me.enable = enable
	return me
}

func (me *hue) Name() string {
	return "hue"
}
