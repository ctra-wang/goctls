package config

import (
	"errors"
	"strings"
)

// DefaultFormat defines a default naming style
const DefaultFormat = "go_zero"

// DefaultGoZeroVersion defines the default version of go zero for migrate
const DefaultGoZeroVersion = "v1.5.4"

// DefaultToolVersion defines the default version of simple admin tools for migrate
const DefaultToolVersion = "v1.5.13"

// GoctlsVersion is goctls version
const GoctlsVersion = "v1.5.16"

// LangEnvKey is the environment variable name to control the tools help info language
const LangEnvKey = "SIMPLE_ADMIN_TOOLS_LANG"

// Config defines the file naming style
type Config struct {
	// NamingFormat is used to define the naming format of the generated file name.
	// just like time formatting, you can specify the formatting style through the
	// two format characters go, and zero. for example: snake format you can
	// define as go_zero, camel case format you can it is defined as goZero,
	// and even split characters can be specified, such as go#zero. in theory,
	// any combination can be used, but the prerequisite must meet the naming conventions
	// of each operating system file name.
	// Note: NamingFormat is based on snake or camel string
	NamingFormat string `yaml:"namingFormat"`
}

// NewConfig creates an instance for Config
func NewConfig(format string) (*Config, error) {
	if len(format) == 0 {
		format = DefaultFormat
	}
	cfg := &Config{NamingFormat: format}
	err := validate(cfg)
	return cfg, err
}

func validate(cfg *Config) error {
	if len(strings.TrimSpace(cfg.NamingFormat)) == 0 {
		return errors.New("missing namingFormat")
	}
	return nil
}
