package utils

import (
	"fmt"
	"reflect"
	"strings"
)

/*
使用反射机制,获取对象每个成员的tag值,用于生成命令行参数列表
对应的成员变量类型只支持三种:
string: 普通参数 如 -async  0
bool:  不带参数值的参数,如  -y
包含flexName,flexValue两个成员变量的结构体数组. 如  -map 0:v  -map 2:a; -c:0:v libx264  -c:a aac
以上三种类型完全可以覆盖fg的所有参数形式
*/
func ReflectByTag(me interface{}) []string {
	args := make([]string, 0)
	elems := reflect.TypeOf(me).Elem()
	values := reflect.ValueOf(me).Elem()
	for i := 0; i < elems.NumField(); i++ {
		field := elems.Field(i)
		tag, ok := field.Tag.Lookup("args") //没包含args标签的忽略
		if !ok {
			continue
		}
		fieldVal := values.FieldByName(field.Name)
		if fieldVal.Kind() == reflect.String {
			fv := fieldVal.String()
			if fv != "" {
				args = append(args, tag)
				args = append(args, fv)
			}
		} else if fieldVal.Kind() == reflect.Bool && fieldVal.Bool() {
			args = append(args, tag)
		} else if fieldVal.Kind() == reflect.Slice { //针对切片类型,需要遍历出每一个元素的成员变量
			for j := 0; j < fieldVal.Len(); j++ {
				sliceValue := fieldVal.Index(j) //切片中的具体元素
				if sliceValue.Kind() != reflect.Struct {
					panic("need struct kind")
				}
				nameArgs := []string{}
				flexName := sliceValue.FieldByName("flexName")            //字段flexName的值
				flexValue := sliceValue.FieldByName("flexValue").String() //flexValue的值
				for ai := 0; ai < flexName.Len(); ai++ {
					itemValue := flexName.Index(ai).String()
					if itemValue != "" {
						nameArgs = append(nameArgs, itemValue)
					}
				}
				argString := tag
				if len(nameArgs) > 0 {
					argString = fmt.Sprintf("%s:%s", argString, strings.Join(nameArgs, ":"))
				}
				args = append(args, strings.TrimRight(argString, ":"), flexValue)
			}
		}
	}
	return args
}
