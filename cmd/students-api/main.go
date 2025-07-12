package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Sankalp-Space/REST-API-project/internal/config"
	"github.com/Sankalp-Space/REST-API-project/internal/http/handlers/student"
)

func main() {
	//load config
	cfg:= config.MustLoad()


	//database connection

	//setup router
	router:=http.NewServeMux()

	router.HandleFunc("POST /api/students",student.New()) // Assuming you have a student package with a New function that returns an http.HandlerFun

	//setup server
	server:=http.Server {
		Addr: cfg.HTTPServer.Addr,
		Handler: router,
	}

	
	slog.Info("starting server", slog.String("address", cfg.Addr))
	fmt.Println("Server is running on", cfg.Addr)
	done:=make(chan os.Signal,1)
	signal.Notify(done, os.Interrupt,syscall.SIGINT,syscall.SIGTERM)
	go func(){
		err:=server.ListenAndServe()
		if err != nil {
		log.Fatal("failed to start the server: ", err)
	}
	}()


	<- done
	slog.Info("shutting down server...")
	ctx, cancel:=context.WithTimeout(context.Background(),5* time.Second)

	defer cancel();
	
	err:=server.Shutdown(ctx);
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown gracefully") 

}
