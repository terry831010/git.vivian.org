package ffmpeg

import (
	"git.vivian.org/ffmpeg/utils"
)

/*
基于命令行对ffprobe调用的封装
*/
type ffprobe struct {
	ffprobePath    string
	loglevel       string `args:"-v"`
	of             string `args:"-of"`
	input          string `args:"-i"`
	format         string `args:"-f"`
	select_streams string `args:"-select_streams"`
	show_streams   bool   `args:"-show_streams"`
	hide_banner    bool   `args:"-hide_banner"`
	show_entries   string `args:"-show_entries"`
}

func FFprobe(ffprobePath string) *ffprobe {
	return &ffprobe{
		ffprobePath: ffprobePath,
	}
}

func (me *ffprobe) Loglevel(level string) *ffprobe {
	me.loglevel = level
	return me
}

func (me *ffprobe) OutputFormat(format string) *ffprobe {
	me.of = format
	return me
}

//输入文件
func (me *ffprobe) Input(in string) *ffprobe {
	me.input = in
	return me
}

//强制指定格式
func (me *ffprobe) Format(format string) *ffprobe {
	me.format = format
	return me
}

//隐藏多余信息
func (me *ffprobe) HideBanner() *ffprobe {
	me.hide_banner = true
	return me
}

//选择流
func (me *ffprobe) SelectStream(select_streams string) *ffprobe {
	me.select_streams = select_streams
	return me
}

//显示流
func (me *ffprobe) ShowStreams() *ffprobe {
	me.show_streams = true
	return me
}

//show_entries
func (me *ffprobe) ShowEntries(show_entries string) *ffprobe {
	me.show_entries = show_entries
	return me
}

//返回命令行参数
func (me *ffprobe) Encode() []string {
	args := utils.ReflectByTag(me)
	return args
}

func (me *ffprobe) BinPath() string {
	return me.ffprobePath
}

//执行
func (me *ffprobe) Exec() (Result, error) {
	return ExecCommand(me)
}
