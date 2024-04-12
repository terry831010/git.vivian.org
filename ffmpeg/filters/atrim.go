package filters

//声音剪切
type atrim struct {
	filter
	start        string `args:"start"`
	end          string `args:"end"`
	start_pts    string `args:"start_pts"`
	end_pts      string `args:"end_pts"`
	duration     string `args:"duration"`
	start_sample string `args:"start_sample"`
	end_sample   string `args:"end_sample"`
}

func ATrim() *atrim {
	c := &atrim{}
	c.meta = c
	return c
}

func (me *atrim) Start(s string) *atrim {
	me.start = s
	return me
}
func (me *atrim) End(e string) *atrim {
	me.end = e
	return me
}
func (me *atrim) StartPts(sp string) *atrim {
	me.start_pts = sp
	return me
}
func (me *atrim) EndPts(ep string) *atrim {
	me.end_pts = ep
	return me
}
func (me *atrim) Duration(d string) *atrim {
	me.duration = d
	return me
}
func (me *atrim) StartSample(start_sample string) *atrim {
	me.start_sample = start_sample
	return me
}
func (me *atrim) EndSample(end_sample string) *atrim {
	me.end_sample = end_sample
	return me
}

func (me *atrim) Name() string {
	return "atrim"
}
