package filters

//纵向翻转
type vflip struct {
	filter
}

func VFlip() *vflip {
	c := &vflip{}
	c.meta = c
	return c
}

func (me *vflip) Name() string {
	return "vflip"
}
