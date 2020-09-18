package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "DWR-932C CLI tool",
		Usage: "Helps manage/control the D-Link DWR-932C LTE Router right from the command line",
		Action: func(c *cli.Context) error {
			fmt.Println("[+] running CLI tool...")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
