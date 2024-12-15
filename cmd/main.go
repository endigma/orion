package main

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:  "orion",
		Usage: "Compile and manage Arma 3 compositions",
		Commands: []*cli.Command{
			bundle,
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "package",
				Aliases: []string{"pkg", "p"},
				Value:   "orion.toml",
				Usage:   "Path to package file",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
				Usage:   "Set the log level",
				Value:   false,
				Action: func(ctx context.Context, c *cli.Command, verbose bool) error {
					if verbose {
						log.SetLevel(log.DebugLevel)
					}

					return nil
				},
			},
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalf("an error occured during compilation: %v", err)
	}
}

func CreateDirectory(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	return nil
}
