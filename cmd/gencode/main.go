package main

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"go-admin/app/helper"
	"go-admin/app/service/haijigit"
	"log"
	"runtime"
	"sync"
)

var (
	wg sync.WaitGroup
)

// 初始化生成所有的代码
func main() {
	maxProcs := runtime.NumCPU()// 获取cpu个数
	runtime.GOMAXPROCS(maxProcs)

	git_chan := make(chan string, 30)
	res_chan := make(chan string, 30)
	//控制最大的进程数
	max_groutine := make(chan bool, 8)


	s_list := haijigit.ServerNameList()
	p_list := haijigit.PackageNameList()

	for _, item := range s_list{
		cmd := "git clone " + gconv.String(item["val"])
		git_chan <- cmd
	}
	for _, item := range p_list {
		cmd := "git clone " + gconv.String(item["val"])
		git_chan <- cmd
	}
	close(git_chan)


	//开8个goroutine
	for git_url := range git_chan{
		cmd := "cd ./code && " + git_url
		wg.Add(1)
		max_groutine <- true
		go func(cmd string, out chan string) {
			wg.Done()
			res, err := helper.Cmd(cmd, true)
			if err != nil {
				log.Println(err)
			}
			out <- gconv.String(res)
			<- max_groutine
		}(cmd, res_chan)
	}

	close(res_chan)

	wg.Wait()
	for str := range res_chan {
		fmt.Println(str)
	}




}
