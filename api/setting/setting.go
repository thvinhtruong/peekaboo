package setting

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerHost        string `mapstructure:"SERVER_HOST"`
	Port              string `mapstructure:"SERVER_PORT"`
	ReadTimeout       string `mapstructure:"SERVER_READ_TIMEOUT"`
	ReadHeaderTimeout string `mapstructure:"SERVER_READ_HEADER_TIMEOUT"`
	WriteTimeout      string `mapstructure:"SERVER_WRITE_TIMEOUT"`
	IdleTimeout       string `mapstructure:"SERVER_IDLE_TIMEOUT"`
	MaxHeaderBytes    string `mapstructure:"SERVER_MAX_HEADER_BYTES"`

	Type         string `mapstructure:"DATABASE_TYPE"`
	User         string `mapstructure:"DATABASE_USER"`
	Password     string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseHost string `mapstructure:"DATABASE_HOST"`
	DatabasePort string `mapstructure:"DATABASE_PORT"`
	Name         string `mapstructure:"DATABASE_NAME"`
	SSLMode      string `mapstructure:"DATABASE_SSL_MODE"`
	CACERTBASE64 string `mapstructure:"DATABASE_CACERTBASE64"`

	CookieDomain   string `mapstructure:"COOKIE_DOMAIN"`
	CookieHttpOnly bool   `mapstructure:"COOKIE_HTTP_ONLY"`
	CookieSecure   bool   `mapstructure:"COOKIE_SECURE"`

	APIDomain string `mapstructure:"API_DOMAIN"`
}

func ReadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
