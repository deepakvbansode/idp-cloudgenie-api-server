package httpserver

import (
	"fmt"
	"net/http"

	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/config"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/ports"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/usecases"
)

type Server struct {
	logger ports.Logger
	Config 		 	*config.WebServerConfig
	Router           *Router
	BlueprintService *usecases.BlueprintService
	ResourceService  *usecases.ResourceService
}

func NewServer(logger ports.Logger, config *config.WebServerConfig, blueprintService *usecases.BlueprintService, resourceService *usecases.ResourceService) *Server {
	r := NewRouter(logger, config, blueprintService, resourceService)
	r.InitializeRouter()
	return &Server{
		logger:          logger,
		Config: 		 config,
		Router:           r,
		
	}
}

func (s *Server) Start() error {
	// Setup routes
	fmt.Printf("Running server on port:", s.Config.Port)
	err := http.ListenAndServe(":"+s.Config.Port, s.Router)
	if err != nil {
		panic(err)
	}
	return nil
}
