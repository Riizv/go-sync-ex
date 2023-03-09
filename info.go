package main

import (
	"fmt"
	"time"
	"runtime"
	"net"
	"log"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err!= nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main(){
	fmt.Println("Informations about you: ")
	fmt.Println("Operating system: " + runtime.GOOS)
	fmt.Println("Architecture: " + runtime.GOARCH)
	fmt.Println("IP: " + GetOutboundIP().String())
	time.Sleep(3 * time.Second)


}