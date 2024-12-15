package config_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"
	"github.com/endigma/orion/config"
)

func TestPackageRead(t *testing.T) {
	want := &config.Package{
		Name:   "Test Package",
		Author: "Test Author",
		Compositions: []config.Composition{
			{
				Name:     "Test Composition",
				Version:  "0.1.0",
				Category: "Test Category",
				InitFile: "testdata/init.sqf",
			},
		},
	}

	got, err := config.ReadPackageFile("testdata/orion.toml")

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}
