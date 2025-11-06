package main

import (
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/adaptors/driven/crossplane"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/adaptors/driven/github"
	mongodb "github.com/deepakvbansode/idp-cloudgenie-backend/internal/adaptors/driven/mongo"
	httpserver "github.com/deepakvbansode/idp-cloudgenie-backend/internal/adaptors/driver/http-server"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/common/logger"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/config"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/usecases"
)

func main() {
	
	config, err := config.GetEnvConfig()
	if err != nil {
		panic("No config")
	}
	logger := logger.NewLogger(config.LogLevel)
	logger.Info("Starting CloudGenie Backend Service", nil)
	crossplaneAdaptor := crossplane.NewCrossplaneAdaptor(logger, config.Crossplane)
	if err != nil {
		logger.Panic("Failed to initialize Crossplane adaptor: %v", err)
	}
	blueprintService := usecases.NewBlueprintService(logger,crossplaneAdaptor) // Pass actual CrossplanePort implementation
	
	githubAdaptor:= github.NewGithubAdaptor(logger, config.Github)	
	if err != nil {
		logger.Panic("Failed to initialize Github adaptor: %v", err)
	}
	mongoRepository := mongodb.NewRepositoryAdaptor(logger,config.Mongo)
	if err != nil {
		logger.Panic("Failed to initialize MongoDB repository: %v", err)
	}
	resourceService := usecases.NewResourceService(logger,githubAdaptor, mongoRepository )   // Pass actual ResourcePort implementation
	server := httpserver.NewServer(logger, config, blueprintService, resourceService)
	server.Start()
}