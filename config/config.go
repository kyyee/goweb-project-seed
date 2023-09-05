package config

import "goweb-project-seed/framework/core/xconfig"

type AppConfig struct {
	Name      string             `mapstrucure:"name"`
	Bootstrap *xconfig.Bootstrap `mapstrucure:"bootstrap"`
}
