package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hanifmhilmy/proj-dompet-api/app/delivery"
	"github.com/hanifmhilmy/proj-dompet-api/app/delivery/middlewares"
	"github.com/hanifmhilmy/proj-dompet-api/app/registry"
	"github.com/hanifmhilmy/proj-dompet-api/config"
	"github.com/julienschmidt/httprouter"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

const (
	port = "API_PORT"
)

func main() {
	cnf, err := config.InitConfig()
	if err != nil {
		log.Panic("[Failed to initialize] Config is not set properly please check! err -> ", err)
	}

	// init the apps container
	ctn, err := registry.NewContainer(cnf)
	if err != nil {
		log.Panic("[Failed to initialize] Module failed to init err -> ", err)
	}
	defer ctn.Clean()

	router := httprouter.New()
	registerHTTPRoute(router, ctn)

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:" + os.Getenv(port),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on port :", os.Getenv(port))
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	// graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	log.Println("Stopping the http server")
	if err := srv.Shutdown(ctx); err != nil {
		log.Println(err)
	}
}

func registerHTTPRoute(r *httprouter.Router, ctn registry.DIContainer) {
	handler := delivery.NewHandler(ctn)

	// initialize middleware for function
	chainMiddleware := func(h http.HandlerFunc, m ...middleware) http.HandlerFunc {
		if len(m) < 1 {
			return h
		}
		wrapped := h

		// loop in reverse to preserve middleware order
		for i := len(m) - 1; i >= 0; i-- {
			wrapped = m[i](wrapped)
		}
		return ctn.HTTPMiddleware(wrapped)
	}

	r.HandlerFunc("GET", "/ping", chainMiddleware(handler.Ping, middlewares.PanicRecoveryMiddleware, middlewares.SetHeaderOptions))
}
