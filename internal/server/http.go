package server

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"

    "example.com/sysinfo/internal/info"
)

// New zwraca wstępnie skonfigurowany http.Server
func New(addr string, sys *info.Info) *http.Server {
    mux := http.NewServeMux()

    // Root – prosty tekst
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    })

    // /api/info – JSON
    mux.HandleFunc("/api/info", func(w http.ResponseWriter, r *http.Request) {
        js, _ := json.MarshalIndent(sys, "", "  ")
        w.Header().Set("Content-Type", "application/json")
        w.Write(js)
    })

    return &http.Server{
        Addr:         addr,
        Handler:      mux,
        ReadTimeout:  5 * time.Second,
        WriteTimeout: 10 * time.Second,
        IdleTimeout:  120 * time.Second,
    }
}
