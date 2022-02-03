package booter

import (
	"context"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/jaronnie/grpc-gateway-example/util"
)

// Running is Running
func Running(bin string) (run bool, pid string) {
	output, _ := exec.Command("pgrep", "-f", bin).Output()
	pidStr := strings.TrimSpace(string(output))

	return !(pidStr == ""), pidStr
}

// run App
func runApp(console bool, a app) (err error) {
	cmd := exec.Command(a.bin, a.args...)

	if console {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return cmd.Run()
	}

	logOutput, err := util.Open(a.log)
	if err != nil {
		return err
	}
	defer logOutput.Close()
	cmd.Stdout = logOutput
	cmd.Stderr = logOutput
	return cmd.Start()
}

// Started is Started
func Started(bin string) (s bool, pid string) {
	// add time out
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
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
