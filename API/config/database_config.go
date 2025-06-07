package config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func LoadDBConfig() *DBConfig {
	return &DBConfig{
		Host:     "localhost",
		Port:     "3306",
		User:     "user",
		Password: "pass",
		Name:     "appDB",
	}
}
