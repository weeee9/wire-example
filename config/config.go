package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database database
	Server   server
}

func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)

	return cfg, err
}

type server struct {
	Host string `envconfig:"APP_SERVER_HOST" default:"localhost"`
	Port string `envconfig:"APP_SERVER_HOST" default:":8080"`
}

type database struct {
	Host     string `envconfig:"APP_DATABASE_HOST" default:"localhost"`
	Port     string `envconfig:"APP_DATABASE_PORT" default:"5432"`
	User     string `envconfig:"APP_DATABASE_USER" default:"user"`
	Password string `envconfig:"APP_DATABASE_PASSWORD" default:"password"`
	Name     string `envconfig:"APP_DATABASE_NAME" default:"database"`
}

func (db database) ConnStr() string {
	connFormat := "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	return fmt.Sprintf(connFormat,
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
	)
}
