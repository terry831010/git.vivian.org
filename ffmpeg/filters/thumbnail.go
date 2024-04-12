package filters

//生成缩略图
type thumbnail struct {
	filter
	n string `args:"n"`
}

func Thumbnail() *thumbnail {
	c := &thumbnail{}
	c.meta = c
	return c
}

func (me *thumbnail) SetN(n string) *thumbnail {
	me.n = n
	return me
}

func (me *thumbnail) Name() string {
	return "thumbnail"
}
