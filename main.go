package main

import (
	"fmt"
	"log"
	"os"

	"github.com/msterzhang/onelist/api"
	"github.com/msterzhang/onelist/initconfig"
	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "onelist",
		Usage: "一个类似emby的专注于刮削alist聚合网盘形成影视媒体库的程序。",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "run, r",
				Usage: "首先运行onelist -run config初始化项目(谨慎操作,否则会覆盖你已有配置文件),会生成config.env配置文件,按要求修改完毕后,运行onelist -run server便可以启动项目,onelist -run admin可查询管理员账户及密码。",
			},
		},
		Action: func(c *cli.Context) error {
			run := c.String("run")
			if run == "config" {
				err := initconfig.InitConfigEnv()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("初始化成功!")
				fmt.Println("修改完config.env配置文件后,运行onelist -run server便可启动项目,忘记密码运行onelist -run admin可查看管理员账户!")
			} else if run == "admin" {
				user, err := initconfig.AdminData()
				if err != nil {
					log.Fatal(err)
				}
				data := fmt.Sprintf("账号:%s 密码:%s", user.UserEmail, user.UserPassword)
				fmt.Println(data)
			} else {
				api.Run()
			}
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
