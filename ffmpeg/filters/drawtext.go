package filters

//视频中写文字
type drawtext struct {
	filter
	fontfile       string `args:"fontfile"`
	text           string `args:"text"`
	textfile       string `args:"textfile"`
	fontcolor      string `args:"fontcolor"`
	fontcolor_expr string `args:"fontcolor_expr"`
	boxcolor       string `args:"boxcolor"`
	bordercolor    string `args:"bordercolor"`
	shadowcolor    string `args:"shadowcolor"`
	box            string `args:"box"`
	boxborderw     string `args:"boxborderw"`
	line_spacing   string `args:"line_spacing"`
	fontsize       string `args:"fontsize"`
	x              string `args:"x"`
	y              string `args:"y"`

	shadowx string `args:"shadowx"`
	shadowy string `args:"shadowy"`
	borderw string `args:"borderw"`
	tabsize string `args:"tabsize"`

	basetime      string `args:"basetime"`
	font          string `args:"font"`
	expansion     string `args:"expansion"`
	timecode      string `args:"timecode"`
	tc24hmax      string `args:"tc24hmax"`
	timecode_rate string `args:"timecode_rate"`

	reload        string `args:"reload"`
	alpha         string `args:"alpha"`
	fix_bounds    string `args:"fix_bounds"`
	start_number  string `args:"start_number"`
	text_shaping  string `args:"text_shaping"`
	ft_load_flags string `args:"ft_load_flags"`
	enable        string `args:"enable"`
}

func DrawText() *drawtext {
	c := &drawtext{}
	c.meta = c
	return c
}

func (me *drawtext) Enable(enable string) *drawtext {
	me.enable = enable
	return me
}

func (me *drawtext) X(x string) *drawtext {
	me.x = x
	return me
}

func (me *drawtext) Y(y string) *drawtext {
	me.y = y
	return me
}

func (me *drawtext) Alpha(alpha string) *drawtext {
	me.alpha = alpha
	return me
}

func (me *drawtext) FontFile(fontfile string) *drawtext {
	me.fontfile = fontfile
	return me
}

func (me *drawtext) Text(text string) *drawtext {
	me.text = text
	return me
}

func (me *drawtext) TextFile(textfile string) *drawtext {
	me.textfile = textfile
	return me
}

func (me *drawtext) FontColor(fontcolor string) *drawtext {
	me.fontcolor = fontcolor
	return me
}
func (me *drawtext) FontColorExpr(fontcolor_expr string) *drawtext {
	me.fontcolor_expr = fontcolor_expr
	return me
}
func (me *drawtext) BoxColor(boxcolor string) *drawtext {
	me.boxcolor = boxcolor
	return me
}
func (me *drawtext) BorderColor(bordercolor string) *drawtext {
	me.bordercolor = bordercolor
	return me
}
func (me *drawtext) ShadowColor(shadowcolor string) *drawtext {
	me.shadowcolor = shadowcolor
	return me
}
func (me *drawtext) Box(box string) *drawtext {
	me.box = box
	return me
}
func (me *drawtext) Boxborderw(boxborderw string) *drawtext {
	me.boxborderw = boxborderw
	return me
}
func (me *drawtext) LineSpacing(line_spacing string) *drawtext {
	me.line_spacing = line_spacing
	return me
}
func (me *drawtext) FontSize(fontsize string) *drawtext {
	me.fontsize = fontsize
	return me
}

func (me *drawtext) Shadowx(shadowx string) *drawtext {
	me.shadowx = shadowx
	return me
}
func (me *drawtext) Shadowy(shadowy string) *drawtext {
	me.shadowy = shadowy
	return me
}

func (me *drawtext) Borderw(borderw string) *drawtext {
	me.borderw = borderw
	return me
}
func (me *drawtext) Tabsize(tabsize string) *drawtext {
	me.tabsize = tabsize
	return me
}
func (me *drawtext) BaseTime(basetime string) *drawtext {
	me.basetime = basetime
	return me
}
func (me *drawtext) Font(font string) *drawtext {
	me.font = font
	return me
}
func (me *drawtext) Expansion(expansion string) *drawtext {
	me.expansion = expansion
	return me
}
func (me *drawtext) TimeCode(timecode string) *drawtext {
	me.timecode = timecode
	return me
}
func (me *drawtext) Tc24hmax(tc24hmax string) *drawtext {
	me.tc24hmax = tc24hmax
	return me
}

func (me *drawtext) TimeCodeRate(timecode_rate string) *drawtext {
	me.timecode_rate = timecode_rate
	return me
}
func (me *drawtext) Reload(reload string) *drawtext {
	me.reload = reload
	return me
}
func (me *drawtext) FixBounds(fix_bounds string) *drawtext {
	me.fix_bounds = fix_bounds
	return me
}
func (me *drawtext) StartNumber(start_number string) *drawtext {
	me.start_number = start_number
	return me
}
func (me *drawtext) TextShaping(text_shaping string) *drawtext {
	me.text_shaping = text_shaping
	return me
}

func (me *drawtext) FtLoadFlags(ft_load_flags string) *drawtext {
	me.ft_load_flags = ft_load_flags
	return me
}

func (me *drawtext) Name() string {
	return "drawtext"
}
