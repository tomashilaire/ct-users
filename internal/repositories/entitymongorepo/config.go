package entitymongorepo

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config interface {
	Dsn() string
	DbName() string
}

type config struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort int
	dbName string
	dsn    string
}

func NewConfig() Config {
	var cfg config
	cfg.dbUser = os.Getenv("DB_USER")
	cfg.dbPass = os.Getenv("DB_PASS")
	cfg.dbHost = os.Getenv("DB_HOST")
	cfg.dbName = os.Getenv("DB_NAME")

	var err error
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}

	cfg.dsn = fmt.Sprintf("mongodb+srv://%s:%s@%s/", cfg.dbUser, cfg.dbPass, cfg.dbHost)
	return &cfg
}

func (c *config) Dsn() string {
	return c.dsn
}

func (c *config) DbName() string {
	return c.dbName
}
