package filters

//改变透明度
type lut struct {
	filter
	c0     string `args:"c0"`
	c1     string `args:"c1"`
	c2     string `args:"c2"`
	c3     string `args:"c3"`
	y      string `args:"y"`
	u      string `args:"u"`
	v      string `args:"v"`
	r      string `args:"r"`
	g      string `args:"g"`
	b      string `args:"b"`
	a      string `args:"a"`
	enable string `args:"enable"`
}

func Lut() *lut {
	c := &lut{}
	c.meta = c
	return c
}

func (me *lut) C0(c0 string) *lut {
	me.c0 = c0
	return me
}

func (me *lut) C1(c1 string) *lut {
	me.c1 = c1
	return me
}

func (me *lut) C2(c2 string) *lut {
	me.c2 = c2
	return me
}

func (me *lut) C3(c3 string) *lut {
	me.c3 = c3
	return me
}

func (me *lut) Y(y string) *lut {
	me.y = y
	return me
}
func (me *lut) U(u string) *lut {
	me.u = u
	return me
}
func (me *lut) V(v string) *lut {
	me.v = v
	return me
}
func (me *lut) R(r string) *lut {
	me.r = r
	return me
}
func (me *lut) G(g string) *lut {
	me.g = g
	return me
}
func (me *lut) B(b string) *lut {
	me.b = b
	return me
}
func (me *lut) A(a string) *lut {
	me.a = a
	return me
}

func (me *lut) Timeline(enable string) *lut {
	me.enable = enable
	return me
}

func (me *lut) Name() string {
	return "lut"
}
