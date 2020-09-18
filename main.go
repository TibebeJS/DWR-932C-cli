package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	var mode string

	app := &cli.App{
		Name:  "DWR-932C CLI tool",
		Usage: "Helps manage/control the D-Link DWR-932C LTE Router right from the command line",
		Commands: []cli.Command{
			{
				Name:    "change-mode",
				Aliases: []string{"m"},
				Usage:   "change connection mode to.",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "to",
						Value:       "auto",
						Usage:       "desired connection mode ('auto', 'LTE', '3G' or '2G').",
						Destination: &mode,
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("[+] running CLI tool...")
					if mode == "auto" {
						fmt.Println("[+] changing connection mode to 'Auto'.")
					} else if mode == "LTE" {
						fmt.Println("[+] changing connection mode to 'LTE'.")
					} else if mode == "3G" {
						fmt.Println("[+] changing connection mode to '3G'.")
					} else if mode == "2G" {
						fmt.Println("[+] changing connection mode to '2G'.")
					}
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
