package config

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Package struct {
	Name   string
	Author string

	Compositions []Composition
}

type Composition struct {
	Name     string
	Version  string
	Category string
	InitFile string
}

func ReadPackageFile(filepath string) (*Package, error) {
	pkgfile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to open package file: %w", err)
	}

	return ReadPackageToml(pkgfile)
}

func ReadPackageToml(bytes []byte) (*Package, error) {
	var pkg Package

	err := toml.Unmarshal(bytes, &pkg)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal package file: %w", err)
	}

	if pkg.Author == "" {
		pkg.Author = "Unknown"
	}

	if pkg.Name == "" {
		return nil, fmt.Errorf("package name is required")
	}

	if len(pkg.Compositions) == 0 {
		return nil, fmt.Errorf("at least one composition is required")
	}

	return &pkg, nil
}
