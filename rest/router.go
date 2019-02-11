package rest

import (
	"net/http"
	"time"

	"github.com/mandocaesar/go-skeleton/common/utility"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mandocaesar/go-skeleton/config"
	"go.elastic.co/apm/module/apmgin"
)

//Router : Instance struct for router model
type Router struct {
	config *config.Configuration
	log    *utility.Log
}

//NewRouter : Instantiate new Router
func NewRouter(configuration *config.Configuration, log *utility.Log) *Router {
	return &Router{config: configuration, log: log}
}

//SetupRouter : register end point
func (r *Router) SetupRouter() *gin.Engine {
	router := gin.New()
	r.log.LogInfo("starting gin")

	//middleware setup

	//APM-gin configuration
	router.Use(apmgin.Middleware(router))

	//CORS-gin configuration
	//TODO : move to yml config
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	v1 := router.Group("api/v1/")
	{
		diagnostic := v1.Group("diagnostic")
		{
			diagnostic.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"Name":       "Kata-Rest-API",
					"message":    "OK",
					"serverTime": time.Now().UTC(),
					"version":    "0.1",
				})
			})
		}
	}

	return router
}
