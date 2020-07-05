package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
)

func GetServerCreds() credentials.TransportCredentials {
	basepath := "/home/tiger/go/src/zhuhui.com/microkit/cert/twoway/"
	cert, err := tls.LoadX509KeyPair(basepath+"server.pem", basepath+"server.key")
	if err != nil {
		log.Fatal(err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(basepath + "ca.pem")
	if err != nil {
		log.Fatal(err)
	}
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	})
	return creds
}

func GetClientCreds() credentials.TransportCredentials {
	basepath := "/home/tiger/go/src/zhuhui.com/microkit/cert/twoway/"
	cert, _ := tls.LoadX509KeyPair(basepath+"client.pem", basepath+"client.key")
	certPool := x509.NewCertPool()
	ca, _ := ioutil.ReadFile(basepath + "ca.pem")
	certPool.AppendCertsFromPEM(ca)
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
		RootCAs:      certPool,
	})
	return creds
}
