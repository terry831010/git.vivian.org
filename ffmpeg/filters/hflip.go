package filters

//横向翻转
type hflip struct {
	filter
}

func HFlip() *hflip {
	c := &hflip{}
	c.meta = c
	return c
}

func (me *hflip) Name() string {
	return "hflip"
}
