package filters

//对音频延迟
type adelay struct {
	filter
	delays string `args:"delays"`
	all    string `args:"all"`
	enable string `args:"enable"`
}

func ADelay() *adelay {
	c := &adelay{}
	c.meta = c
	return c
}

func (me *adelay) Delays(delays string) *adelay {
	me.delays = delays
	return me
}

func (me *adelay) All(a string) *adelay {
	me.all = a
	return me
}

func (me *adelay) Enable(enable string) *adelay {
	me.enable = enable
	return me
}

func (me *adelay) Name() string {
	return "adelay"
}
