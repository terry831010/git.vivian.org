package filters

//读取一个视频作为输入源 与 -i参数功能基本一致
type movie struct {
	filter
	filename      string `args:"filename"`
	format_name   string `args:"f"`
	stream_index  string `args:"si"`
	seek_point    string `args:"sp"`
	streams       string `args:"s"`
	loop          string `args:"loop"`
	discontinuity string `args:"discontinuity"`
}

func Movie() *movie {
	c := &movie{}
	c.meta = c
	return c
}

func (me *movie) Filename(filename string) *movie {
	me.filename = filename
	return me
}
func (me *movie) FormatName(format_name string) *movie {
	me.format_name = format_name
	return me
}
func (me *movie) StreamIndex(stream_index string) *movie {
	me.stream_index = stream_index
	return me
}
func (me *movie) SeekPoint(seek_point string) *movie {
	me.seek_point = seek_point
	return me
}
func (me *movie) Streams(streams string) *movie {
	me.streams = streams
	return me
}
func (me *movie) Loop(loop string) *movie {
	me.loop = loop
	return me
}
func (me *movie) Discontinuity(discontinuity string) *movie {
	me.discontinuity = discontinuity
	return me
}

func (me *movie) Name() string {
	return "movie"
}
