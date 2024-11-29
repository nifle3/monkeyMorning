package config

import "github.com/ilyakaznacheev/cleanenv"

type DB struct {
}

type AI struct {
	Token       string  `env:"TOKEN" env-description:"Token to access "`
	Temperature float64 `env:"TEMPERATURE" env-description:"Model temperature"`
}

type TG struct {
	Token  string `env:"TOKEN" env-description:"token to access your bot. Check BOT_FATHER TELEGRAM"`
	ChatID int64  `env:"CHAT_ID" env-description:"Chat ID to send messages"`
}

type FileStorage struct {
	Static string `env:"STATIC" env-description:"Path to folder with static files"`
}

type Logger struct {
	FilePath string `env:"PATH" env-description:"file for log, must be json"`
	Level    string `env:"LEVEL" env-description:"must be 'warn', 'debug', 'info', 'error', case not sensitive"`
}

type Config struct {
	EnvType string `env:"ENV_TYPE" env-default:"dev" env-description:"must be prod or dev"`
	Logger  `env-prefix:"LOG_"`
	TG      `env-prefix:"TG_"`
	DB
	AI          `env-prefix:"MISTRAL_"`
	FileStorage `env-prefix:"FILE_"`
}

func New() (*Config, error) {
	var result Config
	if err := cleanenv.ReadEnv(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
