package main

import "fmt"
import "os/exec"
import "strings"

func runCommand(cmd string) string {
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err)
	}
	return strings.TrimSuffix(string(out), "\n")
}

func main() {
	status := runCommand("playerctl status")
	if status == "Playing" {
		fmt.Println(runCommand("playerctl metadata artist") + " - " + runCommand("playerctl metadata title"))
	} else {
		fmt.Println("Stopped")
	}
}
