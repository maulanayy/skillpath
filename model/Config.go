package model

type Config struct {
	DB   dbConfig   `yaml:"db"`
	File fileConfig `yaml:"file"`
	Rest restConfig `yaml:"rest"`
}

type dbConfig struct {
	Address  string `yaml:"address"`
	Name     string `yaml:"dbname"`
	Port     int    `yaml:"port"`
	SSL      string `yaml:"ssl"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type DBParams struct {
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	SSLMode  string `json:"ssl_mode"`
}

type restConfig struct {
	Port          string `yaml:"port"`
	TokenValidate bool   `yaml:"token_validate"`
	TokenExpire   int64  `yaml:"token_expire"`
}

type fileConfig struct {
	Path    string `yaml:"path"`
	MaxSize string `yaml:"maxSize"`
}
