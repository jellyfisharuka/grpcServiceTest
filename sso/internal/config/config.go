package config

import "time"

type Config struct {
	Env         string `yaml:"env" env-required:"local"`
	StoragePath string `yaml:"storage_path" env-required:"./data"`
    TokenTTL time.Duration `yaml:"token_ttl" env-required:"true"`
	Grpc GrpcConfig `yaml:"gprc"`
}
type GrpcConfig struct {
	Port    int           `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}
func MustLoad() *Config {
  
}
func fetchConfigPath() string {
	
}
