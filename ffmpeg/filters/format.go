package filters

//修改像素格式
type format struct {
	filter
	pix_fmts string `args:"pix_fmts"`
}

func Format() *format {
	c := &format{}
	c.meta = c
	return c
}

func (me *format) PixFmts(f string) *format {
	me.pix_fmts = f
	return me
}

func (me *format) Name() string {
	return "format"
}
