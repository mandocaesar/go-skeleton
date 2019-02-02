package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-skeleton/config"
	"github.com/go-skeleton/rest"
)

var (
	configuration config.Configuration
	engine        *gin.Engine
	router        *rest.Router
)

func init() {
	cfg, err := config.New("")
}

func main() {

}
