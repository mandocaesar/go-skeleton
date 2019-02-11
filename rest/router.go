package rest

import (
	"errors"
	"net/http"
	"time"

	"go.elastic.co/apm/module/apmgin"

	"github.com/jinzhu/gorm"
	"github.com/mandocaesar/go-skeleton/common/utility"
	"github.com/mandocaesar/go-skeleton/domain/authentication"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/mandocaesar/go-skeleton/config"
)

//Router : Instance struct for router model
type Router struct {
	config         *config.Configuration
	log            *utility.Log
	db             *gorm.DB
	authController *authentication.Controller
	authService    *authentication.Service
}

//NewRouter : Instantiate new Router
func NewRouter(configuration *config.Configuration, log *utility.Log, _db *gorm.DB) (*Router, error) {

	_authService, err := authentication.NewService(_db)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	_authController, err := authentication.NewController(_authService)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return &Router{
		config:         configuration,
		log:            log,
		authController: _authController,
		authService:    _authService,
	}, nil
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

		authentication := v1.Group("authentication")
		{
			authentication.POST("/", r.authController.Login)
		}
	}

	return router
}
