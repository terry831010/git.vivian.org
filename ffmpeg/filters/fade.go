package filters

//淡入淡出
type fade struct {
	filter
	ftype       string `args:"t"`
	start_frame string `args:"s"`
	nb_frames   string `args:"n"`
	alpha       string `args:"alpha"`
	start_time  string `args:"st"`
	duration    string `args:"d"`
	color       string `args:"c"`
	enable      string `args:"enable"`
}

func Fade() *fade {
	c := &fade{}
	c.meta = c
	return c
}

func (me *fade) Type(t string) *fade {
	me.ftype = t
	return me
}

func (me *fade) StartFrame(s string) *fade {
	me.start_frame = s
	return me
}

func (me *fade) NbFrames(n string) *fade {
	me.nb_frames = n
	return me
}
func (me *fade) Alpha(alpha string) *fade {
	me.alpha = alpha
	return me
}
func (me *fade) StartTime(st string) *fade {
	me.start_time = st
	return me
}
func (me *fade) Duration(d string) *fade {
	me.duration = d
	return me
}
func (me *fade) Color(c string) *fade {
	me.color = c
	return me
}

func (me *fade) TimeLine(t string) *fade {
	me.enable = t
	return me
}

func (me *fade) Name() string {
	return "fade"
}
