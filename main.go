package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}

	for _, envKey := range []string{"DWR_URL", "DWR_USERNAME", "DWR_PASSWORD"} {
		if _, ok := os.LookupEnv(envKey); !ok {
			log.Fatal("'" + envKey + "' key must be set")
		}
	}

}

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
					desiredMode, ok := ConnectionModes[mode]

					if !ok {
						log.Fatalln("[-] Not a valid connection mode - only 'auto', 'LTE', '3G' or '2G' are allowed")
					}

					fmt.Println("[+] running CLI tool...")
					fmt.Println("[+] changing connection mode to '" + mode + "' Mode.")
					ChangeMode(desiredMode)
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
