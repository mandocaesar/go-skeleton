package database

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/mandocaesar/go-skeleton/config"
	"github.com/sirupsen/logrus"
)

//Factory : database struct form db factory
type Factory struct {
	config config.Configuration
	log    *logrus.Logger
}

//NewDbFactory : function to generate new database factory
func NewDbFactory(cfg *config.Configuration, logger *logrus.Logger) (*Factory, error) {
	if cfg == nil {
		return nil, errors.New("Error Intantiate new db instance, config is null")
	}
	return &Factory{config: *cfg}, nil
}

//DBConnection : open connection to Database server
func (factory *Factory) DBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(factory.config.Database.DbType, factory.config.Database.ConnectionURI)
	if err != nil {
		return nil, err
	}

	return db, err
}
