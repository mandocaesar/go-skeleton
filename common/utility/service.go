package utility

import (
	"errors"
	"net/http"
	"time"

	health "github.com/hellofresh/health-go"
)

//HealthService struct
type HealthService struct {
	Config string
}

//NewHealthService : instantiate new service of healthcheck
func NewHealthService(config string) (*HealthService, error) {
	return &HealthService{Config: config}, nil
}

//Register : provide http.handler for health check
func (h *HealthService) Register() (http.Handler, error) {
	health.Register(health.Config{
		Name:      "rabbitmq",
		Timeout:   time.Second * 5,
		SkipOnErr: true,
		Check: func() error {
			return errors.New("saya error mas, wik wik wik")
		},
	})

	return health.Handler(), nil
}
