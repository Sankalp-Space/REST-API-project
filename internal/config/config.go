package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env:"HTTP_SERVER_ADDR" env-required:"true"`
}
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"`
	StoragePath string `yaml:"storage_path" env-required:"true"`

	HTTPServer `yaml:"http_server" env-required:"true"`
}

func MustLoad() *Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == ""{
		flags := flag.String("config","", "Path to the configuration file");
		flag.Parse()

		configPath=*flags

		if configPath == "" {
			log.Fatal("CONFIG_PATH is not set")
		}
	}
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist :: %s", configPath)
	}

	var cgf Config

	err:=cleanenv.ReadConfig(configPath, &cgf)

	if err!=nil{
		log.Fatalf("failed to read config file %s", err.Error())
	}

	return &cgf;
}