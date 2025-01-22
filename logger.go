package hyperion

import "log"

// Logger integration for client.
type Logger interface {
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type StdLogger struct{}

func (l *StdLogger) Info(msg string) {
	log.Println("[INFO] " + msg)
}

func (l *StdLogger) Warn(msg string) {
	log.Println("[WARN] " + msg)
}

func (l *StdLogger) Error(msg string) {
	log.Println("[ERROR] " + msg)
}
