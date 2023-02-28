package main

import (
	"bytes"
	"docker-show/routers"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	//utils.Init()
	//fmt.Println(utils.SystemConf.CloudServer.Ip)
	router := routers.InitRouter()
	//定时器应用
	//cronInit()
	port := strconv.Itoa(8888) //utils.SystemConf.LedHttp.Port
	portstr := ":" + port
	err := router.Run(portstr)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
func runCmd(cmdStr string) string {
	list := strings.Split(cmdStr, " ")
	cmd := exec.Command(list[0], list[1:]...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return stderr.String()
	} else {
		return out.String()
	}
}
