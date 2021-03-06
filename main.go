package main

import (
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/joho/godotenv"
	"github.com/manifoldco/promptui"
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
			log.Fatalf("'%s' key not found on the config file", envKey)
		}
	}

}

func main() {
	var mode string
	var turnWIFIOn bool
	var turnWIFIOff bool

	fmt.Println("[+] running CLI tool...")

	app := &cli.App{
		Name:  "DWR-932C CLI tool",
		Usage: "Helps manage/control the D-Link DWR-932C LTE Router right from the command line",
		Commands: []cli.Command{
			{
				Name:    "change-mode",
				Aliases: []string{"m"},
				Usage:   "change connection mode to a desired mode.",
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

					fmt.Printf("[*] changing connection mode to '%s' Mode.\n", mode)
					ChangeMode(desiredMode)
					return nil
				},
			},
			{
				Name:    "mode prompt",
				Aliases: []string{"p"},
				Usage:   "change connection mode to a desired mode (prompt)",
				Flags:   []cli.Flag{},
				Action: func(c *cli.Context) error {
					prompt1 := promptui.Select{
						Label: "Select Mode",
						Items: reflect.ValueOf(ConnectionModes).MapKeys(),
					}

					_, result, er := prompt1.Run()

					if er != nil {
						fmt.Printf("Mode prompt failed %v\n", er)
					}

					desiredMode, _ := ConnectionModes[result]

					fmt.Printf("[*] changing connection mode to '%s' Mode.\n", result)
					ChangeMode(desiredMode)
					return nil
				},
			},
			{
				Name:    "wifi",
				Aliases: []string{"w"},
				Usage:   "manage WIFI related settings",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:        "on",
						Usage:       "Turn WIFI on",
						Destination: &turnWIFIOn,
					},
					&cli.BoolFlag{
						Name:        "off",
						Usage:       "Turn WIFI off",
						Destination: &turnWIFIOff,
					},
				},
				Action: func(c *cli.Context) error {
					if turnWIFIOn && turnWIFIOff {
						log.Fatal("--on and --off flags are mutually exclusive. please choose one")
					}

					if turnWIFIOn {
						fmt.Println("[+] turing WiFi on...")
						DisableWiFi(!turnWIFIOn)
					} else if turnWIFIOff {
						fmt.Println("[+] turing WiFi off...")
						DisableWiFi(turnWIFIOff)
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
