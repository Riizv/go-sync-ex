package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Riizv/go-sync-ex/internal/info"
	"github.com/Riizv/go-sync-ex/internal/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// --- CLI flags ----------------------------------------------------------
	port := flag.String("port", ":8080", "port na którym wystartuje serwer HTTP (np. ':8080')")
	verbose := flag.Bool("verbose", false, "wypisz szczegóły systemu przy starcie")
	debug := flag.Bool("debug", false, "debug which show more informations")
	//silent := flag.Bool("silent", false, "silent mode")
	//coldStart := flag.Bool("coldStart", false, "boot once again")
	//healthCheck := flag.Bool("healthCheck", false, "health")
	flag.Parse()

	//TODO: add new CLI params

	// --- Informacje o systemie ---------------------------------------------
	systemInformation, err := info.CollectBasicInfo()
	if err != nil {
		log.Printf("WARN: %v", err)
	}
	if *verbose {
		fmt.Printf("App built with %s\n", systemInformation.Version)
		fmt.Printf("Operating system: %s\nArchitecture: %s\nShell: %s\n",
			systemInformation.OS, systemInformation.Arch, systemInformation.Shell)
		fmt.Printf("Local IP: %s | Public IP: %s\n", systemInformation.LocalIP, systemInformation.PublicIP)
		fmt.Printf("UUID: %s\n", systemInformation.UUID)
	}

	if *debug {
		// group added for try it, not intended to be used on production
		fmt.Println("Welcome to debug mode!")
		fmt.Println("PID: ", os.Getpid())
		// group, errg := os.Getgroups()
		hostname, errh := os.Hostname()
		if errh != nil {
			fmt.Println("Err durning retrieving hostname:", errh)
			return
		}
		// if errg != nil {
		// 	fmt.Println("Err durning retrieving group", errg)
		// 	return
		// }
		fmt.Println("Hostname:", hostname)
		// fmt.Println("Groups: ", group)
	}

	// --- Konfiguracja serwera ----------------------------------------------
	srv := server.NewServerService(*port, systemInformation)

	// --- Graceful shutdown --------------------------------------------------
	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() { // uruchamiamy asynchronicznie
		if err := srv.ListenAndServe(); err != nil && err != server.ErrSrvClosed {
			log.Fatalf("server error: %v", err)
		}
	}()
	log.Printf("Serwer wystartował na %s", *port)

	<-ctx.Done() // czekamy na sygnał (co ta linijka robi? Sprawdzić.)
	log.Printf("Otrzymano sygnał, zamykam...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("graceful shutdown nie powiódł się: %v", err)
	}
	log.Println("Zamknięto poprawnie – do zobaczenia!")
}
