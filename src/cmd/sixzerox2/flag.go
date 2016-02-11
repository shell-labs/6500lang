package main

import (
	"config/iniconf"
	"fmt"
	"logger"
	"model/mysqlrepo"
	"net/http"
	"server"

	"github.com/codegangsta/cli"
)

var defaultConf = ""

func initApp(app *cli.App) {
	app.Name = "6500lang"
	app.Version = "1.0.0"

	app.Commands = []cli.Command{
		apiServer(),
		initDB(),
		dropDB(),
	}
}

// Web服务器
func apiServer() cli.Command {
	cmd := cli.Command{
		Name:  "run",
		Usage: "运行server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "config, c",
				EnvVar: "APP_CONFIG",
				Value:  defaultConf,
				Usage:  "Configure file path",
			},
		},
		Action: func(c *cli.Context) {
			confPath := c.String("config")
			serve := server.NewServer(confPath)
			fmt.Println(logo)
			fmt.Printf("Run: http://localhost:%d\n", serve.Conf.Server.Port)

			listen := fmt.Sprintf(":%d", serve.Conf.Server.Port)
			if err := http.ListenAndServe(listen, serve.Pipeline.HTTPHandler()); err != nil {
				serve.Logger.Errorf("http.ListenAndServe(%s, server) err: %s", listen, err)
			}
		},
	}

	return cmd
}

func initDB() cli.Command {
	cmd := cli.Command{
		Name:  "init",
		Usage: "初始化",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "config, c",
				EnvVar: "APP_CONFIG",
				Value:  defaultConf,
				Usage:  "Configure file path",
			},
		},
		Action: func(c *cli.Context) {
			confPath := c.String("config")
			conf, _ := iniconf.New(confPath).Load()
			log := logger.NewBeeLogger(conf)

			if repo, err := mysqlrepo.New(conf, log); err != nil {
				log.Error(err)
			} else {
				repo.InitDB()
			}
		},
	}

	return cmd
}

func dropDB() cli.Command {
	cmd := cli.Command{
		Name:  "drop",
		Usage: "删除数据库表",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config, c",
				Value: defaultConf,
				Usage: "Configure file path",
			},
		},
		Action: func(c *cli.Context) {
			confPath := c.String("config")
			conf, _ := iniconf.New(confPath).Load()
			log := logger.NewBeeLogger(conf)

			if repo, err := mysqlrepo.New(conf, log); err != nil {
				log.Error(err)
			} else {
				repo.DropDB()
			}
		},
	}

	return cmd
}
