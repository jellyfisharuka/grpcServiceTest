package config

import (
	"flag"
	"os"
	"time"
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env-required:"local"`
	StoragePath string        `yaml:"storage_path" env-required:"./data"`
	TokenTTL    time.Duration `yaml:"token_ttl" env-required:"true"`
	Grpc        GrpcConfig    `yaml:"gprc"`
}
type GrpcConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("config path is empty")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist" + err.Error())
	}
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config file: " + err.Error())
	}
	return &cfg

}
func fetchConfigPath() string {
	var res string
	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}

//fetchConfigPath fetches config path from command line flag or environment variable.
//Priority: flag>env>default; Default value is empty string
