package main

import (
	"flag"
	"fmt"
	"github.com/json-iterator/go/extra"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func main() {
	extra.RegisterFuzzyDecoders()
	flag.Parse()
	time.LoadLocation("Asia/Shanghai")

	app := cli.NewApp()
	app.Name = "雅思听力语料库练习工具"
	app.Description = "雅思听力语料库练习工具"
	app.Version = "1.0.0"

	// 多个命令，可以指定到 Commands 中
	app.Commands = []*cli.Command{
		{
			Name:        "check",
			Description: "开始批改",
			Aliases:     []string{"s"},
			Usage: `
「王陆语料库」缩写是 wlylk, 拼音的首字母。如果是第 11 章的第 1 个 Test Paper 那就是 wlylk.11.1。
windows 完整命令是：corpus.exe start wlylk.11.1
`,
			Action: func(ctx *cli.Context) (err error) {
				err = Check(ctx)
				if err != nil {
					fmt.Println("检查出错...", err.Error())
				}
				return nil
			},
		},
		{
			Name:  "help",
			Usage: "TODO 这里是帮助....",
			Action: func(c *cli.Context) error {
				fmt.Println("xxxxxx")
				return nil
			},
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "print the version",
			Action: func(c *cli.Context) error {
				fmt.Println(app.Version)
				return nil
			},
		},
	}

	app.HideVersion = true
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error: %v", err)
	}
}
