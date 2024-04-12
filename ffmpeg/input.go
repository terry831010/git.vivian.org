package ffmpeg

type input struct {
	format   string     `args:"-f"`  //指定格式
	codec    []flexArgs `args:"-c"`  //继承codec指定解码器.如: -c:v:0  h264 .  -c:v 视频解码器  -c:a 音频解码器  -c:s 字幕解码器
	ss       string     `args:"-ss"` //开始时间
	to       string     `args:"-to"` //结束时间
	duration string     `args:"-t"`  //持续时间

	fps  string `args:"-r"`    //指定帧率
	vn   bool   `args:"-vn"`   //去掉视频数据
	sn   bool   `args:"-sn"`   //去掉视频字幕
	an   bool   `args:"-an"`   //去掉音频数据
	dn   bool   `args:"-dn"`   //作为输入选项,阻止对文件的所有数据流进行过滤，或为任何输出自动选择或映射
	ac   string `args:"-ac"`   //指定声道数
	ar   string `args:"-ar"`   //指定音频采样率
	safe string `args:"-safe"` //当使用-f参数时，忽略安全性读取txt文件

	stream_loop string `args:"-stream_loop"` //Set number of times input stream shall be looped
	loop        string `args:"-loop"`        //开启循环读取
	itsoffset   string `args:"-itsoffset"`   //Set the input time offset.
	re          bool   `args:"-re"`          //输入的读取速度降低到输入的本地帧速率
	bitexact    bool   `args:"-bitexact"`    //为（de）muxer和（de/en）编码器启用位精确模式
	//请勿调整成员变量的顺序
	file string `args:"-i"` //指定输入文件
}

type IInput *input //用于导出input

func Input(file string) *input {
	return &input{file: file}
}

func (me *input) Format(format string) *input {
	me.format = format
	return me
}

func (me *input) Safe(safe string) *input {
	me.safe = safe
	return me
}

func (me *input) Duration(duration string) *input {
	me.duration = duration
	return me
}

//codec数组前两个是编解码器的标识和索引,最后一个是具体的编解码器名称
func (me *input) Codec(label string, index string, codec string) *input {
	fa := flexArgs{flexValue: codec}
	if label != "" {
		fa.flexName = append(fa.flexName, label)
	}
	if index != "" {
		fa.flexName = append(fa.flexName, index)
	}
	me.codec = append(me.codec, fa)
	return me
}

func (me *input) Start(start string) *input {
	me.ss = start
	return me
}

func (me *input) End(end string) *input {
	me.to = end
	return me
}

func (me *input) Fps(fps string) *input {
	me.fps = fps
	return me
}

func (me *input) VN(vn bool) *input {
	me.vn = vn
	return me
}

func (me *input) SN(sn bool) *input {
	me.sn = sn
	return me
}

func (me *input) AN(an bool) *input {
	me.an = an
	return me
}

func (me *input) DN(dn bool) *input {
	me.dn = dn
	return me
}

func (me *input) AC(ac string) *input {
	me.ac = ac
	return me
}

func (me *input) AR(ar string) *input {
	me.ar = ar
	return me
}

func (me *input) StreamLoop(l string) *input {
	me.stream_loop = l
	return me
}

func (me *input) Loop(l string) *input {
	me.loop = l
	return me
}

func (me *input) ItsOffset(offset string) *input {
	me.itsoffset = offset
	return me
}

func (me *input) Re() *input {
	me.re = true
	return me
}

func (me *input) BitExact() *input {
	me.bitexact = true
	return me
}
