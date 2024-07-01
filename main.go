package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("wrong")
		return
	}

	switch os.Args[1] {
	case "dnf":
		if runtime.GOOS == "windows" {
			// 在 Windows 上使用 PowerShell 按顺序执行两个脚本
			cmd := exec.Command("powershell", "-Command", "cd $env:USERPROFILE\\Desktop; ./WeGame.lnk; ./leigod_acc.lnk")
			if Output, err := cmd.CombinedOutput(); err != nil {
				log.Println(err)
			} else {
				log.Println(string(Output))
			}
		} else if runtime.GOOS == "linux" {
			// 在 Linux 上使用 bash 按顺序执行两个脚本
			cmd := exec.Command("bash", "-c", "echo hello && echo dnf??")
			if Output, err := cmd.CombinedOutput(); err != nil {
				log.Println(err)
			} else {
				log.Println(string(Output))
			}
		} else {
			fmt.Println("Unsupported OS")
		}
	case "lol":
		// 这里可以添加 lol 的处理逻辑
	default:
		fmt.Println("wrong arg")
		return
	}
}
