package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-skeleton/config"
	"github.com/go-skeleton/rest"
)

var (
	configuration config.Configuration
	engine        *gin.Engine
	migrate       bool
	seed          bool
)

func init() {
	// flag.BoolVar(&migrate, "migrate", false, "run db migration")
	// flag.BoolVar(&seed, "migrate", false, "run db seeder")
	// flag.Parse()

	cfg, err := config.New("./")
	if err != nil {
		panic(fmt.Errorf("error parse configuration, reason: %s", err))
	}

	configuration := *cfg
	instance := rest.NewRouter(&configuration)
	engine = instance.SetupRouter()
}

func main() {
	gin.SetMode(configuration.Server.Mode)
	server := &http.Server{
		Addr:    configuration.Server.Addr,
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(fmt.Errorf("Fatal error failed to start rest-api server, reason : %s", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server")

	// give n seconds for server to shutdown gracefully
	duration := time.Duration(configuration.Server.ShutdownTimeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("Failed to shut down server gracefully: %s", err)
	}
	fmt.Printf("Server shutted down")

}
