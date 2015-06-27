package common

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var logger *log.Logger

// Service interface
type Service interface {
	Name() string
	Start() error
	Stop() error
}

// Serve starts a service, and stops it if recieve INT or TERM signal.
func Serve(s Service) {
	signalCh := make(chan os.Signal, 1)
	exitCh := make(chan bool)

	go func() {
		sig := <-signalCh
		logf(s, "recieve signal: %s", sig)
		exitCh <- true
	}()

	// listening INT & TERM signal
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go startService(s, exitCh)

	<-exitCh

	stopService(s)
}

func startService(s Service, exitCh chan<- bool) {
	logf(s, "started")
	err := s.Start()
	if err != nil {
		logf(s, "ended unexpectely: %s", err)
	} else {
		logf(s, "ended")
	}
	exitCh <- true
}

func stopService(s Service) {
	logf(s, "stoping...")
	if err := s.Stop(); err != nil {
		logf(s, "stopped with error: %s", err)
	}
	logf(s, "Bye-bye!")
}

func logf(s Service, format string, args ...interface{}) {
	name := s.Name()
	msg := fmt.Sprintf(format, args...)
	msg = fmt.Sprintf("[Service %s] %s", name, msg)
	if logger != nil {
		logger.Println(msg)
	} else {
		log.Println(msg)
	}
}
