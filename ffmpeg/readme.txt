针对ffmpeg和ffprobe的命令行下的封装【版本：4.4-static】
author：liujinde
2021-12-20

基本思路：
1、观察 ffmpeg的命令行规范:   ffmpeg  [全局参数]  [输入文件与参数列表] [过滤器]  [输出文件与参数列表]
2、将其中的输入、过滤器、输出抽象成三个不同的对象，利用golang中的tag与反射机制将参数与成员变量结合起来。
3、filter所有参数必须是string, ffmpeg  input  output三个核心对象,成员类型有string,bool,[]struct{}{}.
其他类型不予支持,这三种类型通过反射机制已经可以覆盖ffmpeg的所有命令行下的参数形式。
4、最终将类似ffmpeg -y   -i a.mp4  -f overlay=1212;flip=ow=PI/3 b.mp4 的命令封装成以下调用形式：
import (
    "ffmpeg"
    "fmt"
)
ret := ffmpeg.FFmpeg("/usr/bin/ffmpeg").Input(input).Output(output).Filter(filter).Exec()
fmt.Println(ret)
便于在程序中调用。

