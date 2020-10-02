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
	"github.com/hanifmhilmy/proj-dompet-api/app/registry"
	"github.com/hanifmhilmy/proj-dompet-api/config"
	"github.com/julienschmidt/httprouter"
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
		Addr:         "0.0.0.0:1234",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Println("Listening on port :1234")
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
	m := func(h http.HandlerFunc) http.HandlerFunc {
		return ctn.HTTPMiddleware(h)
	}

	r.HandlerFunc("GET", "/ping", m(handler.Ping))
}
