package ffmpeg

import (
	"log"
	"os"
	"strings"

	"git.vivian.org/ffmpeg/filters"
	"git.vivian.org/ffmpeg/utils"
)

/*
基于命令行对ffmpeg4.4调用的封装
*/
type ffmpeg struct {
	ffmpegPath             string
	loglevel               string `args:"-v"` //set logging level
	overwrite              bool   `args:"-y"`
	recast_media           string `args:"-recast_media"`           //Allow forcing a decoder of a different media type than the one detected or designated by the demuxer
	filter_complex_threads string `args:"-filter_complex_threads"` //The default is the number of available CPUs
	stats                  bool   `args:"-stats"`                  //Print encoding progress/statistics. It is on by default
	stats_period           string `args:"-stats_period"`           //Set period at which encoding progress/statistics are updated. Default is 0.5 seconds.
	progress               string `args:"-progress"`               //Send program-friendly progress information to url.
	timelimit              string `args:"-timelimit"`              //Exit after ffmpeg has been running for duration seconds in CPU user time.
	sinks                  string `args:"-sinks"`                  //-sinks pulse,server=192.168.0.4
	sources                string `args:"-sources"`                //-sources pulse,server=192.168.0.4

	report      bool   `args:"-report"`      //上报日志
	hide_banner bool   `args:"-hide_banner"` //隐藏多余信息
	cpuflags    string `args:"-cpuflags"`    //允许设置和清除cpu标志
	max_alloc   string `args:"-max_alloc"`   //设置ffmpeg的malloc函数族在堆上分配块的最大大小限制
	hwaccel     string `args:"-hwaccel"`     //硬件加速方法,如 vdpau,cuda,cuvid

	ignore_unknown bool `args:"-ignore_unknown"` //如果尝试复制未知类型的输入流，请忽略这些流，而不是失败
	copy_unknown   bool `args:"-copy_unknown"`   //如果尝试复制未知类型的输入流，则允许复制这些流，而不是失败

	map_channel string `args:"-map_channel"` //将音频通道从给定输入映射到输出。如果未设置output_file_id.stream_说明符，则音频通道将映射到所有音频流上。
	vsync       string `args:"-vsync"`       //视频同步方法   passthrough 0 ,vfr 2,cfr 1,drop,auto -1
	async       string `args:"-async"`       //Audio sync method. "Stretches/squeezes" the audio stream to match the timestamps

	input  []*input  //支持多个input
	output []*output //支持多个output
	filter []filters.Filter
}

/*
针对 -c:v:0 libx264 -c:a:0 libx264; -map  1:a, -map 0:v  这样类型的参数,统一通过此结构体来接收
*/
type flexArgs struct {
	flexName  []string //可能出现在参数名称后的字符串,比如v,0,以:拼接在主参数名后面. 形如: -c:a:0
	flexValue string   //参数值
}

func FFmpeg(ffmpegPath string) *ffmpeg {
	return &ffmpeg{
		ffmpegPath: ffmpegPath,
		input:      make([]*input, 0),
		output:     make([]*output, 0),
		filter:     make([]filters.Filter, 0),
	}
}

func (me *ffmpeg) Loglevel(level string) *ffmpeg {
	me.loglevel = level
	return me
}

func (me *ffmpeg) OverWrite() *ffmpeg {
	me.overwrite = true
	return me
}

func (me *ffmpeg) Stats() *ffmpeg {
	me.stats = true
	return me
}

func (me *ffmpeg) RecastMedia(rm string) *ffmpeg {
	me.recast_media = rm
	return me
}
func (me *ffmpeg) FilterComplexThreads(n string) *ffmpeg {
	me.filter_complex_threads = n
	return me
}
func (me *ffmpeg) StatsPeriod(p string) *ffmpeg {
	me.stats_period = p
	return me
}
func (me *ffmpeg) Progress(url string) *ffmpeg {
	me.progress = url
	return me
}

func (me *ffmpeg) Timelimit(timelimit string) *ffmpeg {
	me.timelimit = timelimit
	return me
}

//显示输出设备接收器
func (me *ffmpeg) Sinks(s string) *ffmpeg {
	me.sinks = s
	return me
}

//显示输入设备接收器
func (me *ffmpeg) Sources(s string) *ffmpeg {
	me.sources = s
	return me
}

//日志以文件形式保存
func (me *ffmpeg) Report() *ffmpeg {
	me.report = true
	return me
}

//忽略未知编码
func (me *ffmpeg) IgnoreUnknown() *ffmpeg {
	me.ignore_unknown = true
	return me
}

//复制时忽略未知编码
func (me *ffmpeg) CopyUnknown() *ffmpeg {
	me.copy_unknown = true
	return me
}

func (me *ffmpeg) HideBanner() *ffmpeg {
	me.hide_banner = true
	return me
}

//Allows setting and clearing cpu flags
func (me *ffmpeg) CpuFlags(c string) *ffmpeg {
	me.cpuflags = c
	return me
}

//最大内存分配
func (me *ffmpeg) MaxAlloc(m string) *ffmpeg {
	me.max_alloc = m
	return me
}

//视频同步方法
func (me *ffmpeg) VSync(vsync string) *ffmpeg {
	me.vsync = vsync
	return me
}

//音频方法
func (me *ffmpeg) ASync(async string) *ffmpeg {
	me.async = async
	return me
}

func (me *ffmpeg) MapChannel(c string) *ffmpeg {
	me.map_channel = c
	return me
}

//硬件加速
func (me *ffmpeg) HwAccel(h string) *ffmpeg {
	me.hwaccel = h
	return me
}

//支持多个输入
func (me *ffmpeg) Input(in *input) *ffmpeg {
	me.input = append(me.input, in)
	return me
}

//支持多个输出
func (me *ffmpeg) Output(out *output) *ffmpeg {
	me.output = append(me.output, out)
	return me
}

//过滤器,可以设置多个
func (me *ffmpeg) Filter(filter filters.Filter) *ffmpeg {
	me.filter = append(me.filter, filter)
	return me
}

//转成命令行参数列表
func (me *ffmpeg) Encode() []string {
	args := utils.ReflectByTag(me)
	//输入文件参数
	for _, v := range me.input {
		inputArgs := utils.ReflectByTag(v)
		args = append(args, inputArgs...)
		log.Println("input args:", inputArgs)
	}
	//过滤器
	filterArgs := []string{}
	for _, v := range me.filter {
		filterArgs = append(filterArgs, v.Encode())
	}
	if len(filterArgs) > 0 {
		filterArgStr := strings.Join(filterArgs, ";") //形成滤镜图
		log.Println("filter args:", filterArgStr)
		args = append(args, "-filter_complex", filterArgStr)
	}
	//输出文件以及参数
	for _, v := range me.output {
		outputArgs := utils.ReflectByTag(v)
		//输出文件前面的参数
		if len(outputArgs) > 0 {
			args = append(args, outputArgs...)
		}
		args = append(args, v.File())
		log.Println("output args:", outputArgs)
	}
	log.Println("final args:", args)
	return args
}

func (me *ffmpeg) BinPath() string {
	return me.ffmpegPath
}

//执行
func (me *ffmpeg) Exec() (Result, error) {
	result, err := ExecCommand(me)
	//如果执行失败,将out的所有文件都删除掉
	if err != nil {
		for _, out := range me.output {
			f := out.File()
			if f != "" {
				os.Remove(f)
			}
		}
	}
	return result, err
}
