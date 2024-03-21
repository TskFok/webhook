package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webhook/app/process"
	"webhook/bootstrap"
	"webhook/router"
)

func main() {
	//守护进程
	args := os.Args

	if len(args) != 1 && args[1] == "bg" {
		process.InitProcess()
	}

	bootstrap.Init()

	//加载router
	router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 5817),
		Handler:        router.Handle,
		ReadTimeout:    time.Duration(20) * time.Second,
		WriteTimeout:   time.Duration(20) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		//不使用https
		if err := s.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
		//使用https
		//if err := s.ListenAndServeTLS(global.TlsCert, global.TlsKey); err != nil && errors.Is(err, http.ErrServerClosed) {
		//	log.Printf("listen: %s\n", err)
		//}
	}()

	//接收信号关闭
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
