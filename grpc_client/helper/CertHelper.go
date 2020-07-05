package helper

import (
	"crypto/tls"
	"crypto/x509"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
)

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
