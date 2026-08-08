package config

import (
	"os"

	"github.com/pelletier/go-toml"
	"github.com/twodigitss/apio/internal/core/shared"
)

// var defaultPath string = shared.ExpandPath("~/Projects/apio/cmd/tea/config.toml")
var defaultPath string = shared.ExpandPath("~/.config/apio/config.toml")

type Config struct {
	UI     UIConfig
	Colors ColorsConfig
	Path   string
}

type UIConfig struct {
	Glyphs  bool
	Sidebar string
	Borders bool
}

type ColorsConfig struct {
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

func Fallback[T comparable](value, fallback T) T {
	var zero T
	if value == zero {
		return fallback
	}
	return value
}

func Default() Config {
	raw := readRawCfg()

	return Config{
		Path: Fallback(raw.Path, defaultPath),
		UI: UIConfig{
			// ponytail: *bool needed — false is zero value so Fallback can't detect "absent"
			Glyphs:  fallbackBool(raw.UI.Glyphs, true),
			Sidebar: Fallback(raw.UI.Sidebar, "left"),
			Borders: fallbackBool(raw.UI.Borders, true),
		},
		Colors: ColorsConfig{
			GET:     Fallback(raw.Colors.GET, "#A8E6CF"),
			POST:    Fallback(raw.Colors.POST, "#FFD3B6"),
			PUT:     Fallback(raw.Colors.PUT, "#A9DEF9"),
			DELETE:  Fallback(raw.Colors.DELETE, "#FFADAD"),
			PATCH:   Fallback(raw.Colors.PATCH, "#D8B4F8"),
			HEAD:    Fallback(raw.Colors.HEAD, "#A0E7E5"),
			OPTIONS: Fallback(raw.Colors.OPTIONS, "#FFC6FF"),
			EXTRA:   Fallback(raw.Colors.EXTRA, "#b3e6a8"),

			TEXT:      Fallback(raw.Colors.TEXT, "#F5F5F5"),
			SUBTEXT:   Fallback(raw.Colors.SUBTEXT, "#CECECE"),
			INFRATEXT: Fallback(raw.Colors.INFRATEXT, "#ABABAB"),

			BORDER:      Fallback(raw.Colors.BORDER, "#F5F5F5"),
			SUBBORDER:   Fallback(raw.Colors.SUBBORDER, "#4A4A4A"),
			INFRABORDER: Fallback(raw.Colors.INFRABORDER, "#2A2A2A"),
		},
	}
}

// rawConfig is only used during parsing so *bool can distinguish absent vs false.
type tomlCfg struct {
	UI struct {
		Glyphs  *bool  `toml:"glyphs"`
		Sidebar string `toml:"sidebar"`
		Borders *bool  `toml:"borders"`
	} `toml:"ui"`
	Colors ColorsConfig `toml:"colors"`
	Path   string       `toml:"path"`
}

func readRawCfg() tomlCfg {
	var raw tomlCfg
	data, err := os.ReadFile(defaultPath)
	if err != nil {
		panic(err)
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
