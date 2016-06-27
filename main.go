package main

import (
	"fmt"
	n "github.com/Cepave/common/net"
	"github.com/toolkits/net"
)

var LocalIps []string

func InitLocalIps() {
	var err error
	LocalIps, err = net.IntranetIP()
	if err != nil {
		fmt.Println("get intranet ip fail:", err)
	}
}

var PublicIps []string

func InitPublicIps() {
	var err error
	PublicIps, err = n.InternetIP()
	if err != nil {
		fmt.Println("get internet ip fail:", err)
	}
}

func main() {
	InitLocalIps()
	InitPublicIps()
	fmt.Println("Local IPs: %v", LocalIps)
	fmt.Println("Public IPs: %v", PublicIps)
}
