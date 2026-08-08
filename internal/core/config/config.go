package config

import (
	"os"

	"github.com/pelletier/go-toml"
	"github.com/twodigitss/apio/internal/shared"
)

// var defaultPath string = shared.ExpandPath("~/.config/apio/config.toml")
var defaultPath string = shared.ExpandPath("~/Projects/apio/cmd/tea/config.toml")

type Config struct {
	UI     UIConfig
	Colors ColorsConfig
	Path   string
}

type UIConfig struct {
	Glyphs  bool
	Sidebar string
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

	newcfg := ReadCfgFile()

	return Config{
		Path: Fallback(newcfg.Path, defaultPath),
		UI: UIConfig{
			Glyphs:  Fallback(newcfg.UI.Glyphs, true),
			Sidebar: Fallback(newcfg.UI.Sidebar, "left"),
		},
		Colors: ColorsConfig{
			GET:     Fallback(newcfg.Colors.GET, "#A8E6CF"),
			POST:    Fallback(newcfg.Colors.POST, "#FFD3B6"),
			PUT:     Fallback(newcfg.Colors.PUT, "#A9DEF9"),
			DELETE:  Fallback(newcfg.Colors.DELETE, "#FFADAD"),
			PATCH:   Fallback(newcfg.Colors.PATCH, "#D8B4F8"),
			HEAD:    Fallback(newcfg.Colors.HEAD, "#A0E7E5"),
			OPTIONS: Fallback(newcfg.Colors.OPTIONS, "#FFC6FF"),
			EXTRA:   Fallback(newcfg.Colors.EXTRA, "#b3e6a8"),

			TEXT:      Fallback(newcfg.Colors.TEXT, "#F5F5F5"),
			SUBTEXT:   Fallback(newcfg.Colors.SUBTEXT, "#CECECE"),
			INFRATEXT: Fallback(newcfg.Colors.INFRATEXT, "#ABABAB"),

			BORDER:      Fallback(newcfg.Colors.BORDER, "#F5F5F5"),
			SUBBORDER:   Fallback(newcfg.Colors.SUBBORDER, "#4A4A4A"),
			INFRABORDER: Fallback(newcfg.Colors.INFRABORDER, "#2A2A2A"),
		},
	}
}

func ReadCfgFile() Config {
	var cfg Config
	path := defaultPath
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
