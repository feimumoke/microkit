package main

import (
	"zhuhui.com/microkit/kit_provider/util"
)

func genkey() {
	util.GenRSAPubAndPri(1024, "./pem")
}

func main() {

}
