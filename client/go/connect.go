package client

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"io/ioutil"
	"log"
)

type ConnParams struct {
	Host      string
	Port      int
	EnableTLS bool
}

func NewConnParams(host string, port int, enableTLS bool) *ConnParams {
	return &ConnParams{Host: host, Port: port, EnableTLS: enableTLS}
}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := ioutil.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair("cert/client-cert.pem", "cert/client-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}

func Connect(params *ConnParams) (client *grpc.ClientConn, err error) {
	serverAddress := fmt.Sprintf("%s:%d", params.Host, params.Port)
	flag.Parse()
	log.Printf("dial server %s, TLS = %t", serverAddress, params.EnableTLS)

	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())

	if params.EnableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Println("cannot load TLS credentials: ", err)
			config := &tls.Config{
				InsecureSkipVerify: true,
			}
			tlsCredentials = credentials.NewTLS(config)
		}

		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	client, err = grpc.Dial(
		serverAddress,
		transportOption,
	)
	if err != nil {
		log.Println("cannot dial server: ", err)
		return client, err
	}

	return client, nil
}
