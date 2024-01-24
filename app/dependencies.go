package app

import (
	"p3ld3v.dev/template/app/domain"
	"p3ld3v.dev/template/app/services"
)

// Dependencies is a struct that holds all the dependencies for the application.
type Dependencies struct {
	Logger      services.Logger
	DbService   services.DbStore
	UserService services.UserService
	Config      domain.Config
}

func NewDependencies(config domain.Config) (*Dependencies, error) {
	logger := services.NewLogger(config.LogLevel)
	dbService, err := services.NewDbService(config.DbConnection, logger)
	dbService.Connect()
	if err != nil {
		return nil, err
	}
	userService := services.NewUserService(dbService, logger)
	return &Dependencies{
		logger,
		dbService,
		userService,
		config,
	}, nil
}
