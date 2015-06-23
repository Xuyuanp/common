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

	logf := func(format string, v ...interface{}) {
		name := s.Name()
		msg := fmt.Sprintf(format, v...)
		msg = fmt.Sprintf("[Service %s] %s", name, msg)
		if logger != nil {
			logger.Println(msg)
		} else {
			log.Println(msg)
		}
	}

	signalCh := make(chan os.Signal, 1)
	exitCh := make(chan bool)

	go func() {
		sig := <-signalCh
		logf("recieve signal: %s", sig)
		exitCh <- true
	}()

	// listening INT & TERM signal
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logf("started")
		err := s.Start()
		if err != nil {
			logf("ended unexpectely: %s", err)
		} else {
			logf("ended")
		}
		exitCh <- true
	}()

	<-exitCh

	logf("stopping...")
	if err := s.Stop(); err != nil {
		logf("stopped with error: %s", err)
	}
	logf("Bye-bye!")
}
