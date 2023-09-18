package main

import (
	"fmt"
	"os/exec"
)

func exec_look_path() {
	f, err := exec.LookPath("ls")
	if err != nil {
		return
	}
	fmt.Println(f)

}
