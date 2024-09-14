package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv         string `mapstructure:"APP_ENV"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	MySQLDBHost    string `mapstructure:"MYSQL_DB_HOST"`
	MySQLDBPort    string `mapstructure:"MYSQL_DB_PORT"`
	MySQLDBUser    string `mapstructure:"MYSQL_DB_USER"`
	MySQLDBPass    string `mapstructure:"MYSQL_DB_PASS"`
	MySQLDBName    string `mapstructure:"MYSQL_DB_NAME"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't load env file", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Can't Unmarshal env file", err)
	}

	if env.AppEnv == "development" {
		log.Println("development mode: On")
	}

	return &env
}
