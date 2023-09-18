package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func myCliTools() {
	// 简单工具
	//simple_cli()

	// 自定义命令
	//command_cli()

	// 带有 flag的客户端，可以使用总 -flag 方式
	//flag_command_cli()

	//带有子命令
	//sub_command_cli()
}

func main() {
	// 执行 - 启动当前进程的可执行文件 - go run main.go 会运行到此处
	exePath, err := os.Executable()
	if err != nil {
		os.Exit(-1)
	}
	fmt.Println(exePath)

	if os.Args[0] == "DD" {
		// fork出的容器进程 - 进行压测
		fmt.Printf("current pid %d", syscall.Getpid())
		fmt.Println()
		cmd := exec.Command("ls", "/home")
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
			fmt.Println("fork stress error")
			fmt.Println(err)
			os.Exit(-1)
		} else {
			fmt.Println("ERROR", err)
			os.Exit(-1)
		}
	}

	cmd := exec.Command(exePath, "DD")
	cmd.SysProcAttr = &syscall.SysProcAttr{}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Println("ERROR", err)
		os.Exit(-1)
	} else {
		// 得到 fork 出来进程映射在外部命名空间的pid
		fmt.Printf("%v\n", cmd.Process.Pid)
	}
	cmd.Process.Wait()

}
