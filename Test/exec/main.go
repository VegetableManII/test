package main

import (
	"os/exec"
	"time"
)

func main() {
	exec.Command("/bin/zsh", "-c", "echo", "hahaha").Run()
	time.Sleep(5 * time.Second)
}
