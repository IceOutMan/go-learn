package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

func command_cli() {
	app := cli.NewApp()
	app.Name = "meiken-test"
	app.Usage = "This Meiken Test app"
	app.Version = "1.0.0"

	app.Before = func(context *cli.Context) error {
		fmt.Println("Before")
		return nil
	}

	app.Action = func(ctx *cli.Context) {
		fmt.Println("Boom")
	}

	app.Commands = []cli.Command{
		{
			Name:    "mk",
			Aliases: []string{"mk"},
			Usage:   "add a task to the list",
			Action: func(c *cli.Context) error {
				fmt.Println("Args: ", c.Args())
				fmt.Println("added task: ", c.Args().First())
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
