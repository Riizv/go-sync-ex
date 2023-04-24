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
type IP struct {
    Query string
}

func getip2() string {
    req, err := http.Get("http://ip-api.com/json/")
    if err != nil {
        return err.Error()
    }
    defer req.Body.Close()

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        return err.Error()
    }

    var ip IP
    json.Unmarshal(body, &ip)

    return ip.Query
}


// func SetUniqueID(){

// }

func WhatOS() string {
	var sysKind string = runtime.GOOS
	var osName string;
	switch sysKind {
		case "windows":
			osName = "Windows"
		case "darwin":
			osName = "MacOS"
		case "linux":
			osName = "Linux"
	}
	return osName
}

func main(){
	fmt.Println("Informations about you: ")
	fmt.Println("Operating system: " + WhatOS())
	fmt.Println("Architecture: " + runtime.GOARCH)
	fmt.Println("Shell:", os.Getenv("SHELL"))
	fmt.Println("IP: " + GetOutboundIP().String())
	fmt.Println("Public IP: " + getip2())
	time.Sleep(3 * time.Second)


}