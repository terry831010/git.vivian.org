package ffmpeg

import (
	"context"
	"log"
	"os/exec"
	"time"
)

type Result struct {
	Response []byte
	ErrMsg   []byte
}

//command接口: ffmpeg、ffprobe都实现了此接口
type Command interface {
	Encode() []string //返回命令的所有参数
	BinPath() string  //命令路径
}

//调用系统方法执行命令行
func ExecCommand(cmder Command) (Result, error) {
	result := Result{Response: make([]byte, 0), ErrMsg: make([]byte, 0)}
	timeout := 120
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	cmdArray := cmder.Encode()
	log.Printf("exec:%s,params:%+v", cmder.BinPath(), cmdArray)
	cmd := exec.CommandContext(ctx, cmder.BinPath(), cmdArray...)
	response, err := cmd.CombinedOutput()
	log.Println(err, string(response))
	result.Response = response
	if err != nil {
		return result, err
	}
	return result, nil
}
