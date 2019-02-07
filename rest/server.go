package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mandocaesar/go-skeleton/common/utility"
	"github.com/mandocaesar/go-skeleton/config"
)

//Server : Rest-API server structure
type Server struct {
	cfg    config.Configuration
	engine *http.Server
	router *Router
}

//New : Rest-API instantiate
func New(config config.Configuration, logger *utility.Log) (*Server, error) {
	_engine := &http.Server{
		Addr: "8080",
	}
	_router := NewRouter(&config, logger)

	return &Server{cfg: config, engine: _engine, router: _router}, nil
}

//Start : run server
func (s *Server) Start() error {
	fmt.Println("REST-API server started...")
	gin.SetMode("Development")
	s.engine.Handler = s.router.SetupRouter()

	return s.engine.ListenAndServe()
}
