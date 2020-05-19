package helper

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

const (
	STATUS_OK = 1
	STATUS_NO = 0
)

func SuccessResponse(ctx *gin.Context, data interface{})  {
	Response(ctx, STATUS_OK, data, "成功")
}

func ErrorResponse(ctx *gin.Context, data interface{}, msg string)  {
	if msg == "" {
		msg = "失败"
	}
	Response(ctx, STATUS_NO, data, msg)
}

func Response(ctx *gin.Context, status int, data interface{}, msg string)  {
	ctx.JSON(http.StatusOK, gin.H{
		"status" : status,
		"data" : data,
		"msg" : msg,
	})
}

// 解析对应的参数
func GetContent(ctx *gin.Context) []byte  {
	byte_str, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	return byte_str
}

//执行命令，返回数据
func Cmd(cmd string, is_shell bool) ([]byte, error)  {
	if is_shell {
		return exec.Command("bash", "-c", cmd).Output()
	} else {
		return exec.Command(cmd).Output()
	}
}

