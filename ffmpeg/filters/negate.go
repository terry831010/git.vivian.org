package filters

//底片效果
type negate struct {
	filter
	alpha  string `args:"negate_alpha"`
	enable string `args:"enable"`
}

func Negate() *negate {
	c := &negate{}
	c.meta = c
	return c
}

func (me *negate) Alpha(alpha string) *negate {
	me.alpha = alpha
	return me
}

func (me *negate) Timeline(enable string) *negate {
	me.enable = enable
	return me
}

func (me *negate) Name() string {
	return "negate"
}
