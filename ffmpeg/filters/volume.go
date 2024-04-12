package filters

//音量
type volume struct {
	filter
	volume            string `args:"volume"`    //音量:如: 1.0
	precision         string `args:"precision"` //精度
	eval              string `args:"eval"`
	replaygain        string `args:"replaygain"`
	replaygain_preamp string `args:"replaygain_preamp"`
	replaygain_noclip string `args:"replaygain_noclip"`
	enable            string `args:"enable"`
}

func Volume() *volume {
	c := &volume{}
	c.meta = c
	return c
}

func (me *volume) Volume(v string) *volume {
	me.volume = v
	return me
}

func (me *volume) Precision(p string) *volume {
	me.precision = p
	return me
}

func (me *volume) Eval(e string) *volume {
	me.eval = e
	return me
}

func (me *volume) Replaygain(r string) *volume {
	me.replaygain = r
	return me
}

func (me *volume) ReplaygainPreamp(r string) *volume {
	me.replaygain_preamp = r
	return me
}

func (me *volume) ReplaygainNoclip(r string) *volume {
	me.replaygain_noclip = r
	return me
}

func (me *volume) Timeline(enable string) *volume {
	me.enable = enable
	return me
}

func (me *volume) Name() string {
	return "volume"
}
