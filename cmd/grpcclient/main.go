package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"os"
	client "root/client/go"
)

func testUploadFile(grpcClient *client.GrpcClient) {
	grpcClient.UploadFile("tmp/test_file.jpg")
}

func testDownloadFile(grpcClient *client.GrpcClient) {
	d1 := grpcClient.DownloadFile("01640ccc-fad7-488a-85fb-f404bde76a95.jpg")
	err := os.WriteFile("tmp/01640ccc-fad7-488a-85fb-f404bde76a95.jpg", d1.Bytes(), 0644)
	if err != nil {
		log.Fatal("error: ", err)
	}
}

func testSignUp(grpcClient *client.GrpcClient) {
	id, _ := grpcClient.SignUp("thilaire",
		"tomas@test.ag",
		"7410",
		"7410",
		"partner")
	log.Println(id)
}
func testSignIn(grpcClient *client.GrpcClient) string {
	_, token, _ := grpcClient.SignIn("tomas@test.ag",
		"7410")
	return token
}

func testAuthenticate(grpcClient *client.GrpcClient, token string) string {
	id, _ := grpcClient.Authenticate(token)
	log.Println(id)
	return id
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

func authMethods() map[string]bool {
	const authServicePath = "/pb.Authentication/"

	return map[string]bool{
		authServicePath + "Authenticate": true,
	}
}

func main() {
	serverAddress := flag.String("address", "localhost:9092", "the server address")
	enableTLS := flag.Bool("tls", false, "enable SSL/TLS")
	flag.Parse()
	log.Printf("dial server %s, TLS = %t", *serverAddress, *enableTLS)

	transportOption := grpc.WithInsecure()

	if *enableTLS {
		tlsCredentials, err := loadTLSCredentials()
		if err != nil {
			log.Fatal("cannot load TLS credentials: ", err)
		}

		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	cc2, err := grpc.Dial(
		*serverAddress,
		transportOption,
	)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	authClient := client.NewGrpcClient(cc2)

	//testSignUp(authClient)
	token := testSignIn(authClient)
	testAuthenticate(authClient, token)
}
