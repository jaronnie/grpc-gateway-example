package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/jaronnie/grpc-gateway-example/internal"
	"github.com/jaronnie/grpc-gateway-example/pkg/booter"
	"github.com/jaronnie/grpc-gateway-example/pkg/version"
	"github.com/jaronnie/grpc-gateway-example/util"
	"github.com/spf13/cobra"
)

var (
	appName      = "grpc-gateway-example" // must be unique
	exitAfterCmd = true
)

var (
	vFlag    bool   // flag to show version info
	nodaemon bool   // flag to tern on nodaemon mode
	workdir  string // working dir of bin
)

// PreBuild build default command for app
func PreBuild(app interface{}, ctx *internal.Context) (err error) {

	// base sub commands
	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().BoolVarP(&vFlag, "version", "v", false, "version info output")
	rootCmd.PersistentFlags().StringVarP(&workdir, "dir", "", ".", "binary working dir")

	// extra sub flags
	startCmd.PersistentFlags().BoolVarP(&nodaemon, "nodaemon", "d", false, "open daemon mode")

	// add commands
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(restartCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(statusCmd)

	// exec cmd
	return execute(app, ctx)
}

func execute(app interface{}, ctx *internal.Context) (err error) {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return err
	}
	// build context
	ctx.AppName = appName

	// set work dir
	if workdir != "." {
		if e, _ := util.PathExists(workdir); !e {
			fmt.Printf("%s does not exists\n", workdir)
			os.Exit(-1)
		}
		abs, err := filepath.Abs(workdir)
		if err != nil {
			fmt.Println("set dir failed: ", err)
			os.Exit(-1)
		}
		_ = os.Chdir(abs)
	}

	// exit application when cmd mode
	if exitAfterCmd {
		os.Exit(0)
	}

	return nil
}

// root command
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: appName + " is a System sub service",
	Long: `
====================================================
Part of the command line is an auxiliary tool, which
is mainly used to start, stop, restart applications 
and child processes, and query application status.
====================================================`,
	Run: func(cmd *cobra.Command, args []string) {
		if vFlag {
			fmt.Println(vStr())
			os.Exit(0)
		}

		_ = cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version info",
	Long:  "Print version information of " + appName,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(vStr())
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start application",
	Long:  "Start application and all sub processes",
	RunE: func(cmd *cobra.Command, args []string) error {
		// run no daemon
		if nodaemon {
			exitAfterCmd = false
		} else {
			// 后台运行app
			return booter.RunAppWithName(appName)
		}

		return nil
	},
}

// restart when daemon mode
var restartCmd = &cobra.Command{
	Use:   "restart",
	Short: "Restart application",
	Long:  "Restart application and all sub processes",
	RunE: func(cmd *cobra.Command, args []string) (err error) {
		if err = booter.StopAppWithName(appName); err != nil {
			return err
		}
		exitAfterCmd = false
		time.Sleep(time.Second)
		if err = booter.RunAppWithName(appName); err != nil {
			return err
		}

		return nil
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop application",
	Long:  "Stop application and all sub processes",
	RunE: func(cmd *cobra.Command, args []string) error {
		return booter.StopAppWithName(appName)
	},
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Query application status",
	Long:  "Query application status and all sub processes' status",
	RunE: func(cmd *cobra.Command, args []string) error {
		return booter.AppStatus(appName)
	},
}

func vStr() (str string) {
	value, err := version.GetString()
	if err != nil {
		fmt.Println(err)
		return
	}
	str = value
	return
}
