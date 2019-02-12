package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/mandocaesar/go-skeleton/common/database"
	"github.com/mandocaesar/go-skeleton/common/utility"

	"github.com/gin-gonic/gin"
	"github.com/mandocaesar/go-skeleton/config"
	_grpc "github.com/mandocaesar/go-skeleton/grpc"
	"github.com/mandocaesar/go-skeleton/rest"
)

var (
	configuration *config.Configuration
	engine        *gin.Engine
	grpcEngine    *_grpc.Server
	httpServer    *http.Server
	migrate       bool
	seed          bool
	log           *utility.Log
)

func main() {
	flag.BoolVar(&migrate, "migrate", false, "run db migration")
	flag.BoolVar(&seed, "seed", false, "run db seeder")
	flag.Parse()

	//setup configuration
	cfg, err := config.New("./")
	if err != nil {
		panic(fmt.Errorf("error parse configuration, reason: %s", err))
	}

	configuration := cfg

	//setup logger
	_log, err := utility.NewLogger(configuration)
	if err != nil {
		panic(fmt.Errorf("error initilize log, reason: %s", err))
	}
	log = _log

	_factory, err := database.NewDbFactory(configuration)
	if err != nil {
		panic(fmt.Errorf("error initilize DB Factory, reason: %s", err))
	}

	_db, err := _factory.DBConnection()
	if err != nil {
		panic(fmt.Errorf("error initilize DB, reason: %s", err))
	}

	if migrate {
		database.Migrate(_db)
	}

	if seed {
		database.Seed(_db)
	}

	//setup REST-API
	instance, err := rest.NewRouter(configuration, log, _db)
	if err != nil {
		panic(fmt.Errorf("error initilize log, reason: %s", err))
	}
	engine = instance.SetupRouter()

	//setup GRPC
	grpcEngine, err = _grpc.New()
	if err != nil {
		panic(fmt.Errorf("error instantiate grpc , reasson: %s", err))
	}

	gin.SetMode(configuration.Server.Mode)
	fmt.Println(configuration.Server.Addr)
	httpServer := &http.Server{
		Addr:    configuration.Server.Addr,
		Handler: engine,
	}

	go func() {
		fmt.Printf("Running http service on %s \n", configuration.Server.Addr)
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
		fmt.Println("starting GRPC server")
		grpcEngine.Instance.Serve(lis)
	}()

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
