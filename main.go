package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/go-skeleton/config"
	"github.com/go-skeleton/rest"
)

var (
	configuration config.Configuration
	engine        *gin.Engine
	router        *rest.Router
	migrate       bool
	seed          bool
)

func init() {
	flag.BoolVar(&migrate, "migrate", false, "run db migration")
	flag.BoolVar(&seed, "migrate", false, "run db seeder")
	flag.Parse()

}

func main() {

}
