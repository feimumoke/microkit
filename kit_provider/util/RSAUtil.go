package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

func GenRSAPubAndPri(bits int, path string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return err
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	priBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: derStream,
	}
	err = ioutil.WriteFile(path+"/private.pem", pem.EncodeToMemory(priBlock), 0644)
	if err != nil {
		return err
	}

	fmt.Println("========= Gen Private Key Success ===========")
	publicKey := &privateKey.PublicKey
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return err
	}
	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	err = ioutil.WriteFile(path+"/public.pem", pem.EncodeToMemory(publicBlock), 0644)
	if err != nil {
		return err
	}
	fmt.Println("======= Gen Public Key Success ======= ")
	return nil
}
