package pangu

import (
	"github.com/urfave/cli/v2"
)

func newShell() (shell *cli.App) {
	shell = cli.NewApp()
	shell.EnableBashCompletion = true
	shell.UseShortOptionHandling = true
	// 定制版本显示，版本号不可改变
	shell.Version = Version

	cli.VersionFlag = &cli.BoolFlag{
		Name:    `version`,
		Aliases: []string{`v`},
		Usage:   `显示应用程序版本信息`,
	}
	// 定制帮助信息
	cli.HelpFlag = &cli.BoolFlag{
		Name:    `help`,
		Aliases: []string{`h`},
		Usage:   `显示所有命令或者帮助信息`,
	}

	return
}