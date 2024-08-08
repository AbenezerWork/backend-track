package bootstrap

import (
	"github.com/spf13/viper"
	"log"
)

type Env struct {
	DBString string `mapstructure:"DB_STRING"`
	Port     string `mapstructure:"PORT"`
	JWTKey   string `mapstructure:"JWTKey"`
}

func NewEnv() *Env {
	env := Env{}
	viper.AddConfigPath(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
