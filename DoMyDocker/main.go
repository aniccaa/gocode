package main

import (
	"awesomeProject/DoMyDocker/container"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

const usage = `DoMydocker is a simple container runtime implementation.
               The purpose of this project to learn how docker works and how to write 
               a docker by ourselves.
               Enjoy it, just for fun.`

func main() {
	app := cli.NewApp()
	app.Name = "mydocker"
	app.Usage = usage

	app.Commands = []*cli.Command{
		initCommand,
		runCommand,
	}

	app.Before = func(context *cli.Context) error {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(os.Stdout)
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

var initCommand = &cli.Command{}
var runCommand = &cli.Command{
	Name: "run",
	Usage: "Create a container with namespace and cgroup limit mydocker run -ti [command]",
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name: "ti",
			Usage: "enable tty",
		},
	},
	/*
		这里是run命令真正执行的函数：
		1. 判断函数是否存在command
		2. 获取用户指定的command
		3. 调用run function去准备启动容器
	 */
	Action: func(context *cli.Context) error {
		if context.NArg() < 1 {
			return fmt.Errorf("Miss container command ")
		}
		cmd := context.Args().Get(0)
		tty := context.Bool("ti")
		Run(cmd, tty)
		return nil
	},
}

func Run(cmd string, tty bool) {
	parent := container.NewParentProcess(tty, cmd)
	if err := parent.Start(); err != nil {
		logrus.Error(err)
	}
	parent.Wait()
	os.Exit(-1)
}