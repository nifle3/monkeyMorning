package config

import "github.com/ilyakaznacheev/cleanenv"

type DBConfig struct {
}

type FileStorageConfig struct {
}

type Config struct {
	DBConfig
	FileStorageConfig
	EnvType string
}

func New() (*Config, error) {
	var result Config
	if err := cleanenv.ReadEnv(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
