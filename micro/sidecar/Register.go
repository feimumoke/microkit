package sidecar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var RegistryURI = "http://localhost:8000"

type JSONRequest struct {
	Jsonrpc string
	Method  string
	Params  []*Service
	Id      int
}

func NewJsonRequest(service *Service, endpoint string) *JSONRequest {
	return &JSONRequest{Jsonrpc: "2.0",
		Method: endpoint,
		Params: []*Service{service},
		Id:     1}
}

func requestRegistry(req *JSONRequest) error {
	b, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
		return err
	}
	resp, err := http.Post(RegistryURI, "application/json", bytes.NewReader(b))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println("regresult:", string(res))
	return nil
}

func RegService(service *Service) error {
	return requestRegistry(NewJsonRequest(service, "Registry.Register"))
}

func UnRegService(service *Service) error {
	return requestRegistry(NewJsonRequest(service, "Registry.Deregister"))
}
