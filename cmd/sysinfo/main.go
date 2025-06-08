package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os/signal"
	"syscall"
	"time"

	"example.com/sysinfo/internal/info"
	"example.com/sysinfo/internal/server"
)

func main() {
	// --- CLI flags ----------------------------------------------------------
	port := flag.String("port", ":8080", "port na którym wystartuje serwer HTTP (np. ':8080')")
	verbose := flag.Bool("verbose", false, "wypisz szczegóły systemu przy starcie")
	flag.Parse()

	//TODO: add new CLI params

	// --- Informacje o systemie ---------------------------------------------
	sysInfo, err := info.Collect()
	if err != nil {
		log.Printf("WARN: %v", err)
	}
	if *verbose {
		fmt.Printf("App built with %s\n", sysInfo.Version)
		fmt.Printf("Operating system: %s\nArchitecture: %s\nShell: %s\n",
			sysInfo.OS, sysInfo.Arch, sysInfo.Shell)
		fmt.Printf("Local IP: %s | Public IP: %s\n", sysInfo.LocalIP, sysInfo.PublicIP)
		fmt.Printf("UUID: %s\n", sysInfo.UUID)
	}

	// --- Konfiguracja serwera ----------------------------------------------
	srv := server.New(*port, sysInfo)

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

	<-ctx.Done() // czekamy na sygnał
	log.Printf("Otrzymano sygnał, zamykam...")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("graceful shutdown nie powiódł się: %v", err)
	}
	log.Println("Zamknięto poprawnie – do zobaczenia!")
}
