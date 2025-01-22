package hyperion

import "time"

// ConnectionType mechanism.
type ConnectionType string

// List of Hyperion connection mechanisms.
const (
	// ConnectTCP       ConnectionType = "TCP"  // TCP Socket
	// ConnectWebSocket ConnectionType = "WS"   // WebSocket
	ConnectHTTP ConnectionType = "HTTP" // HTTP/S
)

// Config for client.
type Config struct {
	VerboseLog bool       `json:"verbose_log"` // Enable verbose logging
	Connection Connection `json:"connection"`
}

// Connection configuration.
type Connection struct {
	Token   string         `json:"token"`
	Type    ConnectionType `json:"type"`
	Host    string         `json:"host"`
	Port    int            `json:"port"`
	SSL     bool           `json:"ssl"`
	Timeout int            `json:"timeout"` // Time to wait for a server response in seconds (default 30 sec)
}

// GetTimeout returns custom or default timeout.
func (c Config) GetTimeout() time.Duration {
	if c.Connection.Timeout > 0 {
		return time.Second * time.Duration(c.Connection.Timeout)
	}

	return time.Second * 30
}
