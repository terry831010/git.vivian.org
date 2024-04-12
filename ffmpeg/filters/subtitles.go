package filters

//添加字幕
type subtitles struct {
	filter
	filename      string `args:"f"`
	original_size string `args:"original_size"`
	fontsdir      string `args:"fontsdir"`
	alpha         string `args:"alpha"`
	charenc       string `args:"charenc"`
	stream_index  string `args:"si"`
	force_style   string `args:"force_style"`
}

func Subtitles() *subtitles {
	c := &subtitles{}
	c.meta = c
	return c
}

func (me *subtitles) Filename(f string) *subtitles {
	me.filename = f
	return me
}
func (me *subtitles) OriginalSize(s string) *subtitles {
	me.original_size = s
	return me
}
func (me *subtitles) FontsDir(fd string) *subtitles {
	me.fontsdir = fd
	return me
}
func (me *subtitles) Alpha(alpha string) *subtitles {
	me.alpha = alpha
	return me
}
func (me *subtitles) Charenc(c string) *subtitles {
	me.charenc = c
	return me
}
func (me *subtitles) StreamIndex(si string) *subtitles {
	me.stream_index = si
	return me
}
func (me *subtitles) ForceStyle(fs string) *subtitles {
	me.force_style = fs
	return me
}

func (me *subtitles) Name() string {
	return "subtitles"
}
