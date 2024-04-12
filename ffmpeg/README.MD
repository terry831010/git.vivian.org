针对ffmpeg和ffprobe的命令行下的封装【版本：4.4-static】
author：liujinde
2021-12-20

基本思路： 
1、观察 ffmpeg的命令行规范:   ffmpeg  [全局参数]  [输入文件与参数列表] [过滤器]  [输出文件与参数列表]，对应本库的形式是:
ffmpeg.FFmpeg("/usr/bin/ffmpeg").Input(input).Output(output).Filter(filter).Exec()

2、将其中的输入、过滤器、输出抽象成三个不同的对象，利用golang中的tag与反射机制将参数与成员变量结合起来。

3、filter所有参数必须是string, ffmpeg  input  output三个核心对象,成员类型有string,bool,[]struct{}{}.

其他类型不予支持,这三种类型通过反射机制已经可以覆盖ffmpeg的所有命令行下的参数形式。

4、最终将类似ffmpeg -y   -i a.mp4  -f overlay=x=100:y=100,boxblur=lr=21;hflip b.mp4 的命令封装成以下调用形式：

其中滤镜链为:overlay=x=100:y=100,boxblur=lr=21     

import (

    "git.wingyouth.org/ffmpeg"

    "git.wingyouth.org/ffmpeg/filters"

    "fmt"

)
fp := ffmpeg.FFmpeg("/usr/bin/ffmpeg").Input(input).Output(output)

headerFilter := filters.Overlay() //滤镜链第一个滤镜

headerFilter.X(100)

headerFilter.Y(100)

filterChain = append(filterChain, overlayFilter)

#滤镜链加入第二个滤镜

boxblurFilter := filters.Boxblur()

boxblurFilter.LumaRadius("21")

headerFilter.Join(boxblurFilter)

fp.Filter(headerFilter)

flipFilter := filters.HFlip()  //加入第二个滤镜

fp.Filter(flipFilter)

ret := fp.Exec()

fmt.Println(ret)

5、filter目录下已经内置了常用的滤镜,可以参考其中滤镜的写法增加你需要的滤镜。

6、本库本质是将ffmpeg的命令行转成可以通过编程来表达的一个基本库，其核心实现是通过反射机制将参数映射到了结构体的tag上。

目前已支持ffmpeg的几乎所有命令行形式，包括滤镜图形式：

ffmpeg -i input.mp4 -vf "scale=320:-1:flags=lanczos,palettegen" -y palette.pngffmpeg -i input.mp4 -i palette.png -filter_complex "scale=320:-1:flags=lanczos[x];[x][1:v]paletteuse" -y output.gif

7、如果您的ffmpeg版本有新参数，可以通过扩展ffmpeg.go来实现你的需求。


