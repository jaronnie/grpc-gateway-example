package booter

import (
	"context"
	"os/exec"
	"strings"
	"time"
)

// Running is Running
func Running(bin string) (run bool, pid string) {
	output, _ := exec.Command("pgrep", "-f", bin).Output()
	pidStr := strings.TrimSpace(string(output))

	return !(pidStr == ""), pidStr
}

// Started is Started
func Started(bin string) (s bool, pid string) {
	// add time out
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	ticker := time.NewTicker(time.Millisecond * 100)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			if r, pid := Running(bin); r {
				return true, pid
			}
		case <-ctx.Done():
			return false, ""
		}
	}
}

// StartAt started At
func StartAt(pid string) (start string) {
	shell := `ps -p ` + pid + ` -o lstart | awk 'NR==2{print $2,$3,$4,$5}'`
	output, _ := exec.Command("/bin/sh", "-c", shell).Output()
	start = strings.TrimSpace(string(output))

	return start
}
