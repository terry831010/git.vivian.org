package filters

//音频转场
type acrossfade struct {
	filter
	nb_samples string `args:"ns"`
	duration   string `args:"d"`
	overlap    string `args:"o"`
	c1         string `args:"c1"`
	c2         string `args:"c2"`
}

func AcrossFade() *acrossfade {
	c := &acrossfade{}
	c.meta = c
	return c
}

func (me *acrossfade) Ns(ns string) *acrossfade {
	me.nb_samples = ns
	return me
}

func (me *acrossfade) Duration(d string) *acrossfade {
	me.duration = d
	return me
}

func (me *acrossfade) Overlap(o string) *acrossfade {
	me.overlap = o
	return me
}

func (me *acrossfade) C1(c1 string) *acrossfade {
	me.c1 = c1
	return me
}

func (me *acrossfade) C2(c2 string) *acrossfade {
	me.c2 = c2
	return me
}

func (me *acrossfade) Name() string {
	return "acrossfade"
}
