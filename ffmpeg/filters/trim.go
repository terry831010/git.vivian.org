package filters

//视频剪切
type trim struct {
	filter
	start       string `args:"start"`
	end         string `args:"end"`
	start_pts   string `args:"start_pts"`
	end_pts     string `args:"end_pts"`
	duration    string `args:"duration"`
	start_frame string `args:"start_frame"`
	end_frame   string `args:"end_frame"`
}

func Trim() *trim {
	c := &trim{}
	c.meta = c
	return c
}

func (me *trim) Start(s string) *trim {
	me.start = s
	return me
}
func (me *trim) End(e string) *trim {
	me.end = e
	return me
}
func (me *trim) StartPts(sp string) *trim {
	me.start_pts = sp
	return me
}
func (me *trim) EndPts(ep string) *trim {
	me.end_pts = ep
	return me
}
func (me *trim) Duration(d string) *trim {
	me.duration = d
	return me
}
func (me *trim) StartFrame(sf string) *trim {
	me.start_frame = sf
	return me
}
func (me *trim) EndFrame(end_frame string) *trim {
	me.end_frame = end_frame
	return me
}

func (me *trim) Name() string {
	return "trim"
}
