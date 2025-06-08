package info

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/google/uuid"
)

// Struct udostępniany przez /api/info
type Info struct {
	OS        string    `json:"os"`
	Version   string    `json:"goVersion"`
	Arch      string    `json:"arch"`
	Shell     string    `json:"shell"`
	LocalIP   string    `json:"localIP"`
	PublicIP  string    `json:"publicIP"`
	UUID      string    `json:"uuid"`
	Timestamp time.Time `json:"timestamp"`
}

func Collect() (*Info, error) {
	local, err := getOutboundIP()
	if err != nil {
		return nil, fmt.Errorf("local IP: %w", err)
	}
	public, _ := fetchPublicIP() // nie krytyczne – ignorujemy błąd
	return &Info{
		OS:        whatOS(),
		Version:   runtime.Version(),
		Arch:      runtime.GOARCH,
		Shell:     getenvOrUnknown("SHELL"),
		LocalIP:   local.String(),
		PublicIP:  public,
		UUID:      uuidNew(),
		Timestamp: time.Now(),
	}, nil
}

/*---------------------------- helpers ------------------------------------*/

func whatOS() string {
	switch runtime.GOOS {
	case "windows":
		return "Windows"
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	default:
		return runtime.GOOS // cokolwiek system zwrócił
	}
}

func getenvOrUnknown(key string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return "unknown"
}

func getOutboundIP() (net.IP, error) {
	d := net.Dialer{Timeout: 2 * time.Second}
	conn, err := d.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP, nil
}

func fetchPublicIP() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(
		ctx, http.MethodGet, "https://api.ipify.org?format=json", nil)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var msg struct {
		IP string `json:"ip"`
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &msg); err != nil {
		return "", err
	}
	return msg.IP, nil
}

// TODO: Add UNIQE UUID and save to file, check before creating new one and hash this file
func uuidNew() string {
	// w razie braku entropii czy problemów z pkg uuid:
	if id, err := uuid.NewRandom(); err == nil {
		return id.String()
	}
	// awaryjny fallback 128-bitów z crypto/rand
	var b [16]byte
	_, _ = rand.Read(b[:])
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
