package xkafka

type Config struct {
	Url        string `json:"url" mapstructure:"url"`
	SASLEnable bool   `json:"sasl_enable" mapstructure:"sasl_enable"`
	User       string `json:"user" mapstructure:"user"`
	Password   string `json:"password" mapstructure:"password"`
	Mechanism  string `json:"mechanism" mapstructure:"mechanism"`
}
