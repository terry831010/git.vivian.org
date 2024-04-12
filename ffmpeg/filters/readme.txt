针对目前项目常用的过滤器进行封装。
1、过滤器参数类型统一使用string****。
2、过滤器中的方法命名规范:segment的parammeters参数中,关于filter的params参数,key必须与filter方法名对应.
对应规则为:如果方法名为驼峰,则key应该为下划线分隔的字符串.如fade过滤器中TimeLine与Type法,对应segment的parameter完整结构如下:
{"filters":[{"name":"fade","params":{"time_line":"AAAABBBB","type":"in"}}]}
3、通过ffmpeg  -h filter=filtername 查看过滤器参数，将其映射到代码中。
4、目前支持 filtergraph 、filterchain语法。如下语法:
ffmpeg  -i a.mp4  -filter_complex fade=t=in:n=100:alpha=true:st=3:d=3,fade=t=in:s=30:n=50:alpha=true:st=3:d=2;[0]fade=t=out:alpha=true:st=3:d=3:enable='between(t,3,6)'
滤镜链:所有滤镜以逗号分隔，解出的帧依次按照顺序进入每一个滤镜进行处理。
滤镜图:滤镜之间形成了一个图,滤镜之间可以与多个滤镜有关系，一个滤镜的输出可以成为多个滤镜的输入，一个滤镜的输入可以来自其他多个滤镜。
5、具体例子可参考concat过滤器图 https://ffmpeg.org/ffmpeg.html#Complex-filtergraphs。