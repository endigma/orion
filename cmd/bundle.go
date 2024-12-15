package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/endigma/orion/compositions"
	"github.com/endigma/orion/config"
	"github.com/urfave/cli/v3"
)

var bundle = &cli.Command{
	Name:  "bundle",
	Usage: "Bundle scripts into compositions",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"out", "o"},
			Value:   "output",
			Usage:   "Output directory relative to the package file",
		},
	},

	Action: func(ctx context.Context, c *cli.Command) error {
		pkgFile := c.String("package")
		outputDir := c.String("output")

		log.Debug("bundling compositions", "pkgfile", pkgFile, "outputDir", outputDir)

		pkg, err := config.ReadPackageFile(pkgFile)
		if err != nil {
			return fmt.Errorf("error reading package file: %w", err)
		}

		log.Debug("loaded package config", "pkg", pkg)

		os.Chdir(filepath.Dir(pkgFile))

		// Compile compositions
		for _, comp := range pkg.Compositions {
			log.Info("bundling composition", "comp", comp.Name, "category", comp.Category, "init", comp.InitFile)

			// Read contents of init file
			initFile, err := os.ReadFile(comp.InitFile)
			if err != nil {
				return fmt.Errorf("error opening init file: %w", err)
			}

			composition := compositions.Composition{
				Name:     comp.Name,
				Author:   pkg.Author,
				Category: comp.Category,
				Init:     string(initFile),
			}

			compDir := filepath.Join(outputDir, url.PathEscape(comp.Name))

			// Create output directory
			err = os.MkdirAll(compDir, os.ModePerm)
			if err != nil {
				return fmt.Errorf("error creating directory: %w", err)
			}

			// Generate header file
			headerFileName := fmt.Sprintf("%s/header.sqe", compDir)
			headerFile, err := os.Create(headerFileName)
			if err != nil {
				return fmt.Errorf("error creating header file: %w", err)
			}

			err = compositions.RenderHeader(headerFile, composition)
			if err != nil {
				return fmt.Errorf("error rendering header file: %w", err)
			}

			// Generate composition file
			compFileName := fmt.Sprintf("%s/composition.sqe", compDir)
			compFile, err := os.Create(compFileName)
			if err != nil {
				return fmt.Errorf("error creating composition file: %w", err)
			}

			err = compositions.RenderComposition(compFile, composition)
			if err != nil {
				return fmt.Errorf("error rendering composition file: %w", err)
			}
		}

		return nil
	},
}
