package ffmpeg

type output struct {
	format string     `args:"-f"` //指定格式
	file   string     //指定输出文件
	codec  []flexArgs `args:"-c"`  //指定编码器  如: -c:v:0  h264 . -c:v 视频编码器  -c:a 音频编码器  -c:s 字幕编码器
	vn     bool       `args:"-vn"` //去掉视频数据
	sn     bool       `args:"-sn"` //去掉视频字幕
	an     bool       `args:"-an"` //去掉音频数据
	dn     bool       `args:"-dn"` //作为输出选项,禁用数据记录，即自动选择或映射任何数据流

	fps      string `args:"-r"`   //指定帧率
	abitrate string `args:"-b:a"` //指定音频码率
	vbitrate string `args:"-b:v"` //指定视频码率
	bitrate  string `args:"-b"`   //指定整体的码率

	ac      string `args:"-ac"`  //指定声道数
	quality string `args:"-q"`   //指定视频质量   0-maxint  数字越大，质量越差
	aq      string `args:"-q:a"` //指定音频质量
	vq      string `args:"-q:v"` //指定视频质量
	ar      string `args:"-ar"`  //指定音频采样率
	crf     string `args:"-crf"` //指定码率因子[0-51]   18-28合理范围.

	s      string `args:"-s"`      //指定视频分辨率 -s 320x240
	aspect string `args:"-aspect"` //纵横比4:3, 16:9 or 1.3333, 1.7777
	preset string `args:"-preset"` //编码速度与质量的平衡 ultrafast、superfast、veryfast、faster、fast、medium、slow、slower、veryslow、placebo

	pixfmt string `args:"-pix_fmt"` //像素格式
	fs     string `args:"-fs"`      //指定输出的文件大小

	timestamp   string     `args:"-timestamp"`   //Set the recording timestamp in the container.
	metadata    string     `args:"-metadata"`    //指定视频元数据  title="my title"
	target      string     `args:"-target"`      //vcd, svcd, dvd, dv, dv50
	dframes     string     `args:"-dframes"`     //Set the number of data frames to output
	vframes     string     `args:"-vframes"`     //Set the number of video frames to output
	attach      string     `args:"-attach"`      //Add an attachment to the output file
	program     string     `args:"-program"`     //创建具有指定标题program_num的程序
	apad        string     `args:"-apad"`        //Pad the output audio stream(s). Argument 是一串过滤器参数，组成与apad过滤器相同
	disposition []flexArgs `args:"-disposition"` //Sets the disposition for a stream  -disposition:a:1  defalut
	shortest    bool       `args:"-shortest"`    //Finish encoding when the shortest input stream ends.
	bitexact    bool       `args:"-bitexact"`    //为（de）muxer和（de/en）编码器启用位精确模式
	maps        []flexArgs `args:"-map"`         //map,通过map设置输出流使用的输入流。 -map 0:a -map 1:v
}

type IOutput *output //用于导出output

func Output(file string) *output {
	return &output{file: file}
}

func (me *output) File() string {
	return me.file
}

//codec数组前两个是编解码器的标识和索引,最后一个是具体的编解码器名称
//label,index可为空
func (me *output) Codec(label string, index string, codec string) *output {
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

func (me *output) Format(format string) *output {
	me.format = format
	return me
}

func (me *output) VN(vn bool) *output {
	me.vn = vn
	return me
}

func (me *output) SN(sn bool) *output {
	me.sn = sn
	return me
}

func (me *output) AN(an bool) *output {
	me.an = an
	return me
}

func (me *output) DN(dn bool) *output {
	me.dn = dn
	return me
}

func (me *output) Fps(fps string) *output {
	me.fps = fps
	return me
}

func (me *output) Abitrate(abitrate string) *output {
	me.abitrate = abitrate
	return me
}

func (me *output) Vbitrate(vbitrate string) *output {
	me.vbitrate = vbitrate
	return me
}

//指定码率
func (me *output) Bitrate(bitrate string) *output {
	me.bitrate = bitrate
	return me
}

func (me *output) AC(ac string) *output {
	me.ac = ac
	return me
}

//视频质量
func (me *output) Quality(q string) *output {
	me.quality = q
	return me
}

//音频质量
func (me *output) AQ(aq string) *output {
	me.aq = aq
	return me
}

//视频质量
func (me *output) VQ(vq string) *output {
	me.vq = vq
	return me
}

func (me *output) AR(ar string) *output {
	me.ar = ar
	return me
}

func (me *output) Crf(crf string) *output {
	me.crf = crf
	return me
}

//视频大小
func (me *output) Size(size string) *output {
	me.s = size
	return me
}

//纵横比
func (me *output) Aspect(aspect string) *output {
	me.aspect = aspect
	return me
}

//编码速度与质量的平衡
func (me *output) Preset(preset string) *output {
	me.preset = preset
	return me
}

//生成map参数
func (me *output) Map(v string) *output {
	fa := flexArgs{flexValue: v}
	me.maps = append(me.maps, fa)
	return me
}

//像素格式
func (me *output) PixFmt(pixfmt string) *output {
	me.pixfmt = pixfmt
	return me
}

//指定输出的文件大小
func (me *output) FileSize(fs string) *output {
	me.fs = fs
	return me
}

//指定文件时间戳
func (me *output) Timestamp(t string) *output {
	me.timestamp = t
	return me
}

//指定输出的文件元数据
func (me *output) Metadata(m string) *output {
	me.metadata = m
	return me
}

//指定输出应用的目标设备
func (me *output) Target(t string) *output {
	me.target = t
	return me
}

//设置数据帧数
func (me *output) Dframes(n string) *output {
	me.dframes = n
	return me
}

//设置视频帧数
func (me *output) Vframes(n string) *output {
	me.vframes = n
	return me
}

//增加附件到输出文件上
func (me *output) Attachment(f string) *output {
	me.attach = f
	return me
}

//指定标题
func (me *output) Program(p string) *output {
	me.program = p
	return me
}

//填充输出音频流
func (me *output) Apad(apad string) *output {
	me.apad = apad
	return me
}

//设置流的处理
func (me *output) Disposition(label string, index string, codec string) *output {
	fa := flexArgs{flexValue: codec}
	if label != "" {
		fa.flexName = append(fa.flexName, label)
	}
	if index != "" {
		fa.flexName = append(fa.flexName, index)
	}
	me.disposition = append(me.disposition, fa)
	return me
}

//当最短的输入流结束时完成编码
func (me *output) Shortest() *output {
	me.shortest = true
	return me
}

func (me *output) BitExact() *output {
	me.bitexact = true
	return me
}
