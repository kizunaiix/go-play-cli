package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"sync"
)

func GoGames(arg string) {

	switch arg {
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
			cmd := exec.Command("bash", "-c", "echo go dnf!")
			if Output, err := cmd.CombinedOutput(); err != nil {
				log.Println(err)
			} else {
				log.Println(string(Output))
			}
		} else {
			fmt.Println("Unsupported OS")
		}

	case "lol":
		if runtime.GOOS == "windows" {
			// 在 Windows 上使用 PowerShell 按顺序执行两个脚本
			cmd := exec.Command("powershell", "-Command", "cd $env:USERPROFILE\\Desktop; ./RiotClientServices.lnk")
			if Output, err := cmd.CombinedOutput(); err != nil {
				log.Println(err)
			} else {
				log.Println(string(Output))
			}
		} else if runtime.GOOS == "linux" {
			// 在 Linux 上使用 bash 按顺序执行两个脚本
			cmd := exec.Command("bash", "-c", "echo go riot!")
			if Output, err := cmd.CombinedOutput(); err != nil {
				log.Println(err)
			} else {
				log.Println(string(Output))
			}
		} else {
			fmt.Println("Unsupported OS")
		}

	case "mi":
		if runtime.GOOS == "windows" {
			// 在 Windows 上使用 PowerShell 按顺序执行两个脚本
			cmd := exec.Command("powershell", "-Command", "cd $env:USERPROFILE\\Desktop; ./mihoyo.lnk")
			if Output, err := cmd.CombinedOutput(); err != nil {
				log.Println(err)
			} else {
				log.Println(string(Output))
			}
		} else if runtime.GOOS == "linux" {
			// 在 Linux 上使用 bash 按顺序执行两个脚本
			cmd := exec.Command("bash", "-c", "echo go mihoyo!")
			if Output, err := cmd.CombinedOutput(); err != nil {
				log.Println(err)
			} else {
				log.Println(string(Output))
			}
		} else {
			fmt.Println("Unsupported OS")
		}
	default:
		fmt.Println("wrong arg")
		return
	}
}

func main() {

	args := os.Args[1:]

	var wg sync.WaitGroup

	for _, arg := range args {
		wg.Add(1)
		argCopy := arg //注意这一行是必须的，这涉及变量的内存分配问题
		go func() {
			defer wg.Done()
			GoGames(argCopy)
		}()
	}
	wg.Wait()
}
