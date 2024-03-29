package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"
	"reflect"
)

func WhatOS() string {
	var sysKind string = runtime.GOOS
	var osName string
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

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

type IP struct {
	Query string
}

func PubIPaddr() string {
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

// func SetUniqueID() string {
//	var uniqueIDsecurePhrase =
// make this shit works
// }


func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func ThreadedHandler() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	time.Sleep(100 * time.Millisecond)
}

func main() {

	fmt.Println("Informations about you: ")
	fmt.Println("The application was built with the Go version: " + runtime.Version())
	fmt.Println("Operating system: " + WhatOS())
	fmt.Println("Architecture: " + runtime.GOARCH)
	fmt.Println("Shell:", os.Getenv("SHELL"))
	fmt.Println("IP: " + GetOutboundIP().String())
	fmt.Println("Public IP: " + PubIPaddr())
	time.Sleep(3 * time.Second)
	var tmt = time.Now()

	fmt.Println(reflect.TypeOf(tmt))
	fmt.Println(reflect.TypeOf(tmt.String()))

	ThreadedHandler()
}
