package filters

//视频模糊
type boxblur struct {
	filter
	luma_radius   string `args:"lr"`
	luma_power    string `args:"lp"`
	chroma_radius string `args:"cr"`
	chroma_power  string `args:"cp"`
	alpha_radius  string `args:"ar"`
	alpha_power   string `args:"ap"`
	timeline      string `args:"enable"`
}

func BoxBlur() *boxblur {
	c := &boxblur{}
	c.meta = c
	return c
}

func (me *boxblur) LumaRadius(lr string) *boxblur {
	me.luma_radius = lr
	return me
}

func (me *boxblur) LumaPower(lp string) *boxblur {
	me.luma_power = lp
	return me
}

func (me *boxblur) ChromaRadius(cr string) *boxblur {
	me.chroma_radius = cr
	return me
}

func (me *boxblur) ChromaPower(cp string) *boxblur {
	me.chroma_power = cp
	return me
}
func (me *boxblur) AlphaRadius(ar string) *boxblur {
	me.alpha_radius = ar
	return me
}
func (me *boxblur) AlphaPower(ap string) *boxblur {
	me.alpha_power = ap
	return me
}

func (me *boxblur) Timeline(enable string) *boxblur {
	me.timeline = enable
	return me
}
func (me *boxblur) Name() string {
	return "boxblur"
}
