package common

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Service interface
type Service interface {
	Start() error
	Stop() error
}

// Serve starts a service, and stops it if recieve INT or TERM signal.
func Serve(s Service) {
	signalCh := make(chan os.Signal, 1)
	exitCh := make(chan bool)

	go func() {
		sig := <-signalCh
		log.Printf("recieve signal: %s", sig)
		exitCh <- true
	}()

	// listening INT & TERM signal
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("start service")
		err := s.Start()
		if err != nil {
			log.Printf("service ended with error: %s", err)
		} else {
			log.Printf("service ended")
		}
		exitCh <- true
	}()

	<-exitCh

	log.Printf("stopping service...")
	if err := s.Stop(); err != nil {
		log.Fatalf("servic ended with error: %s", err)
	}
	log.Printf("Bye-bye!")
}
