package authentication

import (
	"errors"

	"github.com/jinzhu/gorm"
)

//Service for authentication
type Service struct {
	Db *gorm.DB
}

//NewService : instantiate new authentication service
func NewService(db *gorm.DB) (*Service, error) {
	if db == nil {
		return nil, errors.New("failed to intantiate Authentication Service , Db instance is null")
	}
	return &Service{Db: db}, nil
}

//Authenticate : function to authenticate
func (s *Service) Authenticate(input string) string {
	return input
}
