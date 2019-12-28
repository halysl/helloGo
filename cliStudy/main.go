package main

import (
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
	"syscall"
)

func main() {
	app := &cli.App{
		Name: "boom",
		Usage: "make an explosive entrance",
		Description: "you will boom",
		Action: func(c *cli.Context) error {
			fmt.Printf("boom i say, %q", c.Args().Get(0))
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		syscall.Exit(1)
	}
}
