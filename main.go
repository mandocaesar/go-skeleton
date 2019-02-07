package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mandocaesar/go-skeleton/config"
	_grpc "github.com/mandocaesar/go-skeleton/grpc"
	"github.com/mandocaesar/go-skeleton/rest"
)

var (
	configuration config.Configuration
	engine        *gin.Engine
	grpcEngine    *_grpc.Server
	httpServer    *http.Server
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
	grpcEngine, err = _grpc.New()
	if err != nil {
		panic(fmt.Errorf("error instantiate grpc , reasson: %s", err))
	}
}

func ChainProcess() {
	gin.SetMode(configuration.Server.Mode)
	httpServer := &http.Server{
		Addr:    configuration.Server.Addr,
		Handler: engine,
	}

	go func() {
		if err := httpServer.ListenAndServe(); err != nil {
			panic(fmt.Errorf("Fatal error failed to start rest-api server, reason : %s", err))
		}
	}()

	go func() {
		// create a listener on TCP port 7777
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
		if err != nil {
			fmt.Printf("failed to listen: %v", err)
		}
		grpcEngine.Instance.Serve(lis)
	}()
}

func GracefullyShutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down server")

	// give n seconds for server to shutdown gracefully
	duration := time.Duration(configuration.Server.ShutdownTimeout) * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Printf("Failed to shut down server gracefully: %s", err)
	}
	grpcEngine.Instance.GracefulStop()
	fmt.Printf("Server shutted down")
}

func main() {
	ChainProcess()
	GracefullyShutdown()
}
