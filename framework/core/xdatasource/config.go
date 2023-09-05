package xdatasource

type Config struct {
	Debug        bool   `json:"debug" mapstructure:"debug"`
	User         string `json:"user" mapstructure:"user"`
	Password     string `json:"password" mapstructure:"password"`
	Host         string `json:"host" mapstructure:"host"`
	Port         int    `json:"port" mapstructure:"port"`
	DbName       string `json:"db_name" mapstructure:"db_name"`
	MaxIdleConns int    `json:"max_idle_conns" mapstructure:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns" mapstructure:"max_open_conns"`
}
