package config

import (
	"os"

	"github.com/pelletier/go-toml"
	"github.com/twodigitss/apio/internal/core/shared"
)

// var defaultPath string = shared.ExpandPath("~/Projects/apio/cmd/tea/config.toml")
var defaultPath string = shared.ExpandPath("~/.config/apio/config.toml")

type config struct {
	Core   coreConfig
	UI     uIConfig
	Colors colorsConfig
	Path   string
}

type coreConfig struct {
	Timeout string
}

type uIConfig struct {
	Glyphs  bool
	Sidebar string
	Borders bool
}

type colorsConfig struct {
	GET     string
	POST    string
	PUT     string
	DELETE  string
	PATCH   string
	HEAD    string
	OPTIONS string
	EXTRA   string

	TEXT      string
	SUBTEXT   string
	INFRATEXT string

	BORDER      string
	SUBBORDER   string
	INFRABORDER string
}

func fallback[T comparable](value, fallback T) T {
	var zero T
	if value == zero {
		return fallback
	}
	return value
}

func Default() config {
	raw := readRawCfg()

	return config{
		Core: coreConfig{
			Timeout: fallback(raw.Core.Timeout, "10s"),
		},
		Path: fallback(raw.Path, defaultPath),
		UI: uIConfig{
			// ponytail: *bool needed — false is zero value so fallback can't detect "absent"
			Glyphs:  fallbackBool(raw.UI.Glyphs, true),
			Sidebar: fallback(raw.UI.Sidebar, "left"),
			Borders: fallbackBool(raw.UI.Borders, true),
		},
		Colors: colorsConfig{
			GET:     fallback(raw.Colors.GET, "#A8E6CF"),
			POST:    fallback(raw.Colors.POST, "#FFD3B6"),
			PUT:     fallback(raw.Colors.PUT, "#A9DEF9"),
			DELETE:  fallback(raw.Colors.DELETE, "#FFADAD"),
			PATCH:   fallback(raw.Colors.PATCH, "#D8B4F8"),
			HEAD:    fallback(raw.Colors.HEAD, "#A0E7E5"),
			OPTIONS: fallback(raw.Colors.OPTIONS, "#FFC6FF"),
			EXTRA:   fallback(raw.Colors.EXTRA, "#b3e6a8"),

			TEXT:      fallback(raw.Colors.TEXT, "#F5F5F5"),
			SUBTEXT:   fallback(raw.Colors.SUBTEXT, "#CECECE"),
			INFRATEXT: fallback(raw.Colors.INFRATEXT, "#ABABAB"),

			BORDER:      fallback(raw.Colors.BORDER, "#F5F5F5"),
			SUBBORDER:   fallback(raw.Colors.SUBBORDER, "#4A4A4A"),
			INFRABORDER: fallback(raw.Colors.INFRABORDER, "#2A2A2A"),
		},
	}
}

// rawConfig is only used during parsing so *bool can distinguish absent vs false.
type tomlCfg struct {
	Core coreConfig `toml:"core"`
	UI   struct {
		Glyphs  *bool  `toml:"glyphs"`
		Sidebar string `toml:"sidebar"`
		Borders *bool  `toml:"borders"`
	} `toml:"ui"`
	Colors colorsConfig `toml:"colors"`
	Path   string       `toml:"path"`
}

func readRawCfg() tomlCfg {
	var raw tomlCfg
	data, err := os.ReadFile(defaultPath)
	if err != nil {
		return tomlCfg{}
	}
	if err = toml.Unmarshal(data, &raw); err != nil {
		panic(err)
	}
	return raw
}

func fallbackBool(v *bool, def bool) bool {
	if v == nil {
		return def
	}
	return *v
}
