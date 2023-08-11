package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

var language string

func simple_cli() {
	app := cli.NewApp()
	app.Name = "meiken-test"
	app.Usage = "This Meiken Test app"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p", // --port | -p
			Value: 8000,      // default value
			Usage: "listening port",
		},
		cli.StringFlag{
			Name:        "lang, l", // --lang | -l
			Value:       "english", // default english
			Usage:       "read from `FILE`",
			Destination: &language,
		},
	}

	app.Action = func(ctx *cli.Context) {
		fmt.Println("Boom")
		fmt.Println("Global Flags :", ctx.GlobalFlagNames())
		fmt.Println("Global Flag port: ", ctx.GlobalInt("port"))
		fmt.Println("Global Flag lang: ", ctx.GlobalString("lang"))
	}

	app.Before = func(context *cli.Context) error {
		fmt.Println("Before")
		return nil
	}

	app.After = func(c *cli.Context) error {
		fmt.Println("app After")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
