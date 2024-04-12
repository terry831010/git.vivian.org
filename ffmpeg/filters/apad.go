package filters

//填充音频
type apad struct {
	filter
	packet_size string `args:"packet_size"`
	pad_len     string `args:"pad_len"`
	whole_len   string `args:"whole_len"`
	pad_dur     string `args:"pad_dur"`
	whole_dur   string `args:"whole_dur"`
	enable      string `args:"enable"`
}

func APad() *apad {
	c := &apad{}
	c.meta = c
	return c
}
func (me *apad) PacketSize(packet_size string) *apad {
	me.packet_size = packet_size
	return me
}

func (me *apad) PadLen(pad_len string) *apad {
	me.pad_len = pad_len
	return me
}

func (me *apad) WholeLen(whole_len string) *apad {
	me.whole_len = whole_len
	return me
}

func (me *apad) PadDur(pad_dur string) *apad {
	me.pad_dur = pad_dur
	return me
}

func (me *apad) WholeDur(whole_dur string) *apad {
	me.whole_dur = whole_dur
	return me
}

func (me *apad) Enable(enable string) *apad {
	me.enable = enable
	return me
}

func (me *apad) Name() string {
	return "apad"
}
