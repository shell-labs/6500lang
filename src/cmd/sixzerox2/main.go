package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	initApp(app)
	app.Run(os.Args)
}
