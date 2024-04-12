package filters

//音频转场
type afade struct {
	filter
	t            string `args:"t"`
	start_sample string `args:"ss"`
	nb_samples   string `args:"ns"`
	start_time   string `args:"st"`
	duration     string `args:"d"`
	curve        string `args:"c"`
}

func AFade() *afade {
	c := &afade{}
	c.meta = c
	return c
}

func (me *afade) T(t string) *afade {
	me.t = t
	return me
}

func (me *afade) Duration(d string) *afade {
	me.duration = d
	return me
}

func (me *afade) StartSample(ss string) *afade {
	me.start_sample = ss
	return me
}

func (me *afade) NbSamples(ns string) *afade {
	me.nb_samples = ns
	return me
}

func (me *afade) StartTime(ss string) *afade {
	me.start_time = ss
	return me
}

func (me *afade) Curve(c string) *afade {
	me.curve = c
	return me
}

func (me *afade) Name() string {
	return "afade"
}
