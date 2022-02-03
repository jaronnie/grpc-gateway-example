package booter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// RunAppWithName run app with name
func RunAppWithName(bin string) (err error) {
	if r, pid := Running(bin); r {
		fmt.Print("[", bin, "] ", pid, "\n")
		return nil
	}

	// app bin path
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return err
	}

	cmd := exec.Command(filepath.Join(dir, bin), []string{"start", "--nodaemon"}...)

	// no log file default
	if err = cmd.Start(); err != nil {
		fmt.Println("start err")
		return err
	}
	if s, pid := Started(bin); s {
		fmt.Print("[", bin, "] ", pid, "\n")
	}

	return
}

// StopAppWithName stop app with name
func StopAppWithName(bin string) (err error) {
	run, pid := Running(bin)
	if !run {
		fmt.Print("[", bin, "] DOWN\n")
		return nil
	}

	cmd := exec.Command("kill", "-TERM", pid)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err == nil {
		fmt.Print("[", bin, "] DOWN\n")
	}
	return
}

// AppStatus show app status
func AppStatus(bin string) (err error) {
	if r, pid := Running(bin); r {

		fmt.Print("[", bin, "] ", pid, ", Started At: ", StartAt(pid), "\n")
		return nil
	}
	fmt.Print("[", bin, "] DOWN\n")
	return nil
}
