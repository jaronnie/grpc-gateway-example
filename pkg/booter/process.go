package booter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/jaronnie/grpc-gateway-example/util"
	"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type app struct {
	name string
	bin  string
	log  string
	args []string
}

type process struct {
	c    *cobra.Command
	apps []app
}

func (p *process) init() *process {
	// set workdir
	workdir := p.c.Flag("dir").Value.String()
	if workdir != "." {
		if err := os.Chdir(workdir); err != nil {
			fmt.Println(fmt.Println(err))
		}
	}
	// load apps & bins
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// target dir
	target := p.c.Flag("work").Value.String()

	err = filepath.Walk(wd, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() && !strings.HasSuffix(path, ".log") {
			// add to global
			splits, module := filepath.Split(path)

			apn := module
			if list := strings.Split(module, "-"); len(list) == 2 {
				apn = list[1]
			}

			p.apps = append(p.apps, app{
				name: apn,
				bin:  path,
				log:  filepath.Join(splits, module+".log"),
				args: []string{"start", "--nodaemon", "--dir", filepath.Join(target, apn)},
			})
		}
		return nil
	})
	if err != nil {
		return nil
	}

	return p
}

func (p *process) filter(names []string) (aps []app) {
	if len(names) == 0 {
		return p.apps
	}
	aps = []app{}
	for _, a := range p.apps {
		if util.StringInSlice(a.name, names) {
			aps = append(aps, a)
		}
	}

	// all apps
	if len(names) == 1 && names[0] == "all" {
		aps = p.apps
	}

	if len(aps) == 0 {
		// not found
		fmt.Println("app not found: ", names)
		os.Exit(-1)
	}
	return
}

// getProcess return a new process
func getProcess(c *cobra.Command) (p *process) {
	p = &process{c: c}
	return p.init()
}

// console : run in foreground
func (p *process) start(console bool, args []string) (err error) {
	apps := p.filter(args)

	for _, app := range apps {
		// Skip starting if the module is already running
		if r, pid := Running(app.bin); r {
			fmt.Print("[", app.bin, "] ", pid, "\n")
			continue
		}

		if err := runApp(console, app); err != nil {
			fmt.Printf("run [%s] app err: %s", app.name, err.Error())
			continue
		}

		if s, pid := Started(app.bin); s {
			fmt.Print("[", app.name, "] ", pid, "\n")
			continue
		}

		fmt.Printf("[%s] failed to start", app.name)
	}
	return nil
}

func (p *process) restart(args []string) (err error) {
	apps := p.filter(args)

	for _, app := range apps {
		if err := p.stop(args); err != nil {
			fmt.Printf("stop [%s] err: %s\n", app.name, err.Error())
			continue
		}

		time.Sleep(1 * time.Second)

		if err := p.start(false, []string{app.name}); err != nil {
			return err
		}
	}

	return nil
}

func (p *process) stop(args []string) (err error) {
	apps := p.filter(args)

	for _, app := range apps {
		run, pid := Running(app.bin)
		if !run {
			fmt.Print("[", app.name, "] DOWN\n")
			continue
		}

		cmd := exec.Command("kill", "-TERM", pid)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err == nil {
			fmt.Print("[", app.name, "] DOWN\n")
			continue
		} else {
			fmt.Printf("stop [%s] err: %s", app.name, err.Error())
		}
	}

	return nil
}

func (p *process) status(args []string) (err error) {
	apps := p.filter(args)

	var pd [][]string
	for _, app := range apps {
		var row []string
		if r, pid := Running(app.bin); r {
			row = append(row, app.name, "UP", pid, StartAt(pid))
		} else {
			row = append(row, app.name, "DOWN", "-", "-")
		}
		pd = append(pd, row)
	}

	table := tablewriter.NewWriter(os.Stdout)
	// TODO add api health check
	table.SetHeader([]string{"App", "State", "Pid", "StartAt"})
	table.SetBorder(false) // Set Border to false
	table.AppendBulk(pd)   // Add print data
	table.Render()

	return nil
}

func (p *process) tail(args []string) (err error) {
	if len(args) < 1 {
		return errors.New("tail : need app name")
	}

	var tailArgs []string = []string{"-f"}

	apps := p.filter(args)

	for _, app := range apps {
		if e, _ := util.PathExists(app.log); !e {
			fmt.Printf("%s does not exists", app.log)
			continue
		}
		tailArgs = append(tailArgs, app.log)
	}

	cmd := exec.Command("tail", tailArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
