package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Riizv/go-sync-ex/internal/configuration"
	"github.com/Riizv/go-sync-ex/internal/info"
	"github.com/Riizv/go-sync-ex/internal/server"
)

func main() {

	configuration.ConfInit()

	// --- CLI flags ----------------------------------------------------------
	port := flag.String("p", ":8080", "port for http server (eg. def ':8080')")
	verbose := flag.Bool("v", false, "at program start list sys params")
	debug := flag.Bool("d", false, "debug which show more information")
	serverStart := flag.Bool("sS", false, "starting server")
	//silent := flag.Bool("silent", false, "silent mode")
	//coldStart := flag.Bool("coldStart", false, "boot once again")
	//healthCheck := flag.Bool("healthCheck", false, "health")
	flag.Parse()

	//TODO: add new CLI params

	// --- System informations ---------------------------------------------
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

	// --- Debug mode ---------------------------------------------
	if *debug {
		fmt.Println("Welcome to debug mode!")
		fmt.Println("PID: ", os.Getpid())
		hostname, errh := os.Hostname()
		if errh != nil {
			fmt.Println("Err durning retrieving hostname:", errh)
			return
		}
		fmt.Println("Hostname:", hostname)

	}

	if *serverStart {
		// --- Server configuration ----------------------------------------------
		srv := server.NewServerService(*port, systemInformation)

		// --- Graceful shutdown --------------------------------------------------
		ctx, stop := signal.NotifyContext(context.Background(),
			syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		go func() { // uruchamiamy asynchronicznie
			if err := srv.ListenAndServe(); err != nil && !errors.Is(err, server.ErrSrvClosed) {
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

}
