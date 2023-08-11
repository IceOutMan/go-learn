package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func flag_command_cli() {
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
			Name:  "lang, l", // --lang | -l
			Value: "english", // default english
			Usage: "read from `FILE`",
		},
	}

	app.Before = func(context *cli.Context) error {
		fmt.Println("Before")
		return nil
	}

	app.Action = func(ctx *cli.Context) {
		fmt.Println("Boom")
		fmt.Println("Global Flags :", ctx.GlobalFlagNames())
		fmt.Println("Global Flag port: ", ctx.GlobalInt("port"))
		fmt.Println("Global Flag lang: ", ctx.GlobalString("lang"))
	}

	app.Commands = []cli.Command{
		{
			Name:    "mk",
			Aliases: []string{"mk"},
			Usage:   "add a task to the list",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "ti",
					Usage: "enable tty",
				},
				// 这里和 ti 是 或的关系
				cli.IntFlag{
					Name:  "age",
					Value: 10,
					Usage: "enable d",
				},
			},
			Action: func(c *cli.Context) error {
				fmt.Println("MK Command")
				fmt.Println("Args: ", c.Args())
				fmt.Println("my flags: ", c.FlagNames())
				fmt.Println("my flag ti: ", c.Bool("ti"))
				fmt.Println("my flag age: ", c.Int("age"))
				return nil
			},
		},
		{
			Name:    "complete",
			Aliases: []string{"c"},
			Usage:   "complete a task on the list",
			Action: func(c *cli.Context) error {
				fmt.Println("completed task: ", c.Args().First())
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
