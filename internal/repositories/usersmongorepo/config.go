package usersmongorepo

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"root/pkg/apperrors"
	"root/pkg/errors"
	"strconv"
	"time"
)

type Config interface {
	Dsn() string
	DbName() string
	DbTlsConfig() *tls.Config
	DbConnTimeOut() time.Duration
}

type config struct {
	dbUser       string
	dbPass       string
	dbHost       string
	dbPort       int
	dbName       string
	rPref        string
	connTimeout  time.Duration
	queryTimeout time.Duration
	caFilePath   string
	tlsConfig    *tls.Config
	dsn          string
}

func NewConfig() Config {
	var cfg config

	// Path to the AWS CA file
	cfg.caFilePath = "rds-combined-ca-bundle.pem"

	// Timeout operations after N seconds
	cfg.connTimeout = 10
	cfg.queryTimeout = 30
	cfg.dbUser = os.Getenv("DB_USER")
	cfg.dbPass = os.Getenv("DB_PASS")
	cfg.dbHost = os.Getenv("DB_HOST")
	cfg.dbName = os.Getenv("DB_NAME")

	// Which instances to read from
	cfg.rPref = "secondaryPreferred"

	connectionURI := "mongodb://%s:%s@%s:%d/?tls=true&replicaSet=rs0&readpreference=%s&tlsAllowInvalidHostnames=true"

	var err error
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}

	cfg.dsn = fmt.Sprintf(connectionURI, cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort, cfg.rPref)
	//cfg.dsn = "mongodb://root:7410@192.168.1.5:27017/"

	cfg.tlsConfig, err = getCustomTLSConfig(cfg.caFilePath)
	if err != nil {
		log.Fatalf("Failed getting TLS configuration: %v", err)
	}
	return &cfg
}

func getCustomTLSConfig(caFile string) (*tls.Config, error) {
	tlsConfig := new(tls.Config)
	certs, err := ioutil.ReadFile(caFile)

	if err != nil {
		return tlsConfig, err
	}

	tlsConfig.RootCAs = x509.NewCertPool()
	ok := tlsConfig.RootCAs.AppendCertsFromPEM(certs)

	if !ok {
		return tlsConfig, errors.LogError(errors.New(apperrors.Internal,
			nil, "Failed parsing pem file", ""))
	}

	return tlsConfig, nil
}

func (c *config) Dsn() string {
	return c.dsn
}

func (c *config) DbName() string {
	return c.dbName
}

func (c *config) DbTlsConfig() *tls.Config {
	return c.tlsConfig
}

func (c *config) DbConnTimeOut() time.Duration {
	return c.connTimeout
}
