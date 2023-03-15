package main

import (
	"fmt"
	"time"
	"runtime"
	"net/http"
	"log"
	"os"
	"net"
	"io/ioutil"
	"encoding/json"
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



func SetUniqueID(){

}

func main(){
	fmt.Println("Informations about you: ")
	fmt.Println("Operating system: " + runtime.GOOS)
	fmt.Println("Architecture: " + runtime.GOARCH)
	fmt.Println("Shell:", os.Getenv("SHELL"))
	fmt.Println("IP: " + GetOutboundIP().String())
	fmt.Println("Public IP: " + getip2())
	time.Sleep(3 * time.Second)


}