package config

import ("github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	DataBase string
	ConnectString string
	Port string
}

func Get_config() Config{
    var conf Config
    if _, err := toml.DecodeFile("server.conf", &conf); err != nil {
		log.Panic(err)
	}
	return conf
}