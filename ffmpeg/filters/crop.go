package filters

//视频裁剪
type crop struct {
	filter
	w           string `args:"w"`
	h           string `args:"h"`
	x           string `args:"x"`
	y           string `args:"y"`
	keep_aspect string `args:"keep_aspect"`
	exact       string `args:"exact"`
}

func Crop() *crop {
	c := &crop{}
	c.meta = c
	return c
}

func (me *crop) SetW(w string) *crop {
	me.w = w
	return me
}

func (me *crop) SetH(h string) *crop {
	me.h = h
	return me
}

func (me *crop) SetX(x string) *crop {
	me.x = x
	return me
}

func (me *crop) SetY(y string) *crop {
	me.y = y
	return me
}

func (me *crop) KeepAspect(keepAspect string) *crop {
	me.keep_aspect = keepAspect
	return me
}

func (me *crop) Exact(exact string) *crop {
	me.exact = exact
	return me
}

func (me *crop) Name() string {
	return "crop"
}
