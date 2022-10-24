package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/app/interface/persistence/rdbms/sqlconnection"
	"server/app/interface/restful/handler"
	"server/setting"
	"strconv"
	"sync"
	"syscall"
	"time"
)

// operation is a clean up function on shutting down
type operation func(ctx context.Context) error

// gracefulShutdown waits for termination syscalls and doing clean up operations after received it
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// Do the operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}

func init() {
	log.Println("Server Started Successfully")
}

func main() {
	config, err := setting.ReadConfig(".")
	if err != nil {
		panic(err)
	}

	sqlconnection.Init(config)
	routers := handler.Routing()

	setting.SetCookieDomain(config.CookieDomain)
	setting.SetCookieHTTPS(config.CookieHttpOnly)
	setting.SetCookieSecure(config.CookieSecure)

	timeOut, _ := strconv.Atoi(config.ReadTimeout)
	headerTimeOut, _ := strconv.Atoi(config.ReadHeaderTimeout)
	writeTimeOut, _ := strconv.Atoi(config.WriteTimeout)
	idleTimeOut, _ := strconv.Atoi(config.IdleTimeout)
	maxHeaderBytes, _ := strconv.Atoi(config.MaxHeaderBytes)

	s := &http.Server{
		Handler:           routers,
		Addr:              ":8080",
		ReadTimeout:       time.Duration(timeOut) * time.Second,
		ReadHeaderTimeout: time.Duration(headerTimeOut) * time.Second,
		WriteTimeout:      time.Duration(writeTimeOut) * time.Second,
		IdleTimeout:       time.Duration(idleTimeOut) * time.Second,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	s.ListenAndServe()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"database": func(ctx context.Context) error {
			return sqlconnection.DBConn.Close()
		},
		"http-server": func(ctx context.Context) error {
			return s.Close()
		},
		// Add other cleanup operations here
	})

	<-wait

}
