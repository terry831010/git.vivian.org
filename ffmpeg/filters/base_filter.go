package filters

import (
	"fmt"
	"reflect"
	"strings"

	"git.vivian.org/ffmpeg/utils"
)

//过滤器父类
type filter struct {
	input  []string //输入流label 如: [frank]
	output []string //输出流label 如: [terry]
	meta   Filter   //子类的指针,用meta表示真正的filter。本filter仅为父类
	chain  []Filter //此过滤器后面跟着的过滤器链, 最终会形成通过逗号分隔开的的过滤器列表: scale=w=30,overlay=x=30
}

//滤镜接口
type Filter interface {
	Input(string) Filter            //设置滤镜的输入流
	Output(string) Filter           //设置滤镜的输出流
	GetInputs() []string            //返回滤镜的输入流
	GetOutputs() []string           //返回滤镜的输出流
	Join(Filter) Filter             //将其他滤镜加入成为滤镜链
	SetParameter(map[string]string) //设置参数
	Encode() string                 //编码成命令行参数字符串
	Name() string
}

//输入标签如 a,b
func (me *filter) Input(input string) Filter {
	me.input = append(me.input, input)
	return me
}

//输出标签 如 1:a
func (me *filter) Output(output string) Filter {
	me.output = append(me.output, output)
	return me
}

//返回filter的输入流数组
func (me *filter) GetInputs() []string {
	return me.input
}

//返回filter的输出流数组
func (me *filter) GetOutputs() []string {
	return me.output
}

//增加filter链
func (me *filter) Join(f Filter) Filter {
	me.chain = append(me.chain, f)
	return me
}

//将本身的成员变量编码成命令行字符串
func (me *filter) Encode() string {
	filterList := []string{}
	args := []string{}
	name := me.meta.Name()
	filterArgs := utils.ReflectByTag(me.meta)
	for i := 0; i < len(filterArgs); i += 2 {
		item := fmt.Sprintf("%s=%s", filterArgs[i], filterArgs[i+1])
		args = append(args, item)
	}
	filterMain := fmt.Sprintf("%s=%s", name, strings.Join(args, ":"))
	filterList = append(filterList, strings.TrimRight(filterMain, "=")) //剔除最后的=,针对filter没有参数的场景

	//将链里的其他filter也加进来
	for _, v := range me.chain {
		itemArgs := []string{}
		chainValues := utils.ReflectByTag(v)
		for i := 0; i < len(chainValues); i += 2 {
			item := fmt.Sprintf("%s=%s", chainValues[i], chainValues[i+1])
			itemArgs = append(itemArgs, item)
		}
		filterChain := fmt.Sprintf("%s=%s", v.Name(), strings.Join(itemArgs, ":"))
		//以下增加对滤镜链中每个滤镜的输入,输出流的支持
		inStream := []string{}
		for _, in := range v.GetInputs() {
			inStream = append(inStream, fmt.Sprintf("[%s]", in))
		}
		outStream := []string{}
		for _, out := range v.GetOutputs() {
			outStream = append(outStream, fmt.Sprintf("[%s]", out))
		}
		chainItem := fmt.Sprintf("%s%s%s", strings.Join(inStream, ""), strings.TrimRight(filterChain, "="), strings.Join(outStream, ""))
		filterList = append(filterList, chainItem)
	}

	inputs := []string{}
	for _, v := range me.input {
		inputs = append(inputs, fmt.Sprintf("[%s]", v))
	}
	outputs := []string{}
	for _, v := range me.output {
		outputs = append(outputs, fmt.Sprintf("[%s]", v))
	}
	//支持这样的格式 :ffmpeg -vf  [0]scale=a=1:b=2,movie=file=****[out];[out]overlay=x=122
	return fmt.Sprintf("%s%s%s", strings.Join(inputs, ""), strings.Join(filterList, ","), strings.Join(outputs, ""))
}

func (me *filter) Name() string {
	return ""
}

/*
根据每个filter提供的接口对应的parameter中的key,通过反射机制,动态调用这些方法,将parameter参数传给filter
规则是如果方法名是驼峰,则以_分隔. 如:atrim有一个方法是EndPts,则对应的parameter的key是:end_pts.
否则就是方法名的小写
*/
func (me *filter) SetParameter(parameter map[string]string) {
	if len(parameter) < 1 {
		return
	}
	obj := reflect.ValueOf(me.meta)
	for k, v := range parameter {
		methodAll := []string{}
		nameArray := strings.Split(k, "_")
		for _, s := range nameArray {
			methodAll = append(methodAll, strings.ToUpper(s[:1])+s[1:])
		}
		methodName := strings.Join(methodAll, "")
		method := obj.MethodByName(methodName)
		zv := reflect.Value{}
		if method != zv {
			method.Call([]reflect.Value{reflect.ValueOf(v)})
		}
	}
}
