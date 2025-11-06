package httpserver

import (
	"net/http"

	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/adaptors/driver/http-server/handlers"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/adaptors/driver/http-server/middlewares"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/config"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/ports"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/usecases"
	"github.com/gorilla/mux"
)
type Router struct {
	logger ports.Logger
	*mux.Router
	Config 		 *config.WebServerConfig
	BlueprintService *usecases.BlueprintService
	ResourceService  *usecases.ResourceService
}

func NewRouter(logger ports.Logger, config *config.WebServerConfig, blueprintService *usecases.BlueprintService, resourceService *usecases.ResourceService) *Router {
	return &Router{
		logger: logger,
		Config: 		 config,
		Router: mux.NewRouter(),
		BlueprintService: blueprintService,
		ResourceService: resourceService,
	}
}

func (r *Router) InitializeRouter() {
	r.initializeMiddleware()
	s := (*r).PathPrefix(r.Config.RoutePrefix).Subrouter()

	s.HandleFunc(HealthCheckURI, handlers.HealthCheckHandler(r.logger)).Methods(http.MethodGet).Name(HealthCheckAPIName)

	s.HandleFunc(BlueprintURI, handlers.GetBlueprintsHandler(r.logger,r.BlueprintService)).Methods(http.MethodGet).Name(GetBlueprintAPIName)

	s.HandleFunc(ResourcesURI, handlers.GetResourcesHandler(r.logger, r.ResourceService)).Methods(http.MethodGet).Name(GetResourceAPIName)
	s.HandleFunc(ResourceDetailURI, handlers.GetResourceHandler(r.logger, r.ResourceService)).Methods(http.MethodGet).Name(GetResourceDetailsAPIName)
	s.HandleFunc(ResourcesURI, handlers.CreateResourceHandler(r.logger, r.ResourceService)).Methods(http.MethodPost).Name(CreateResourceAPIName)
	s.HandleFunc(ResourceDetailURI, handlers.DeleteResourceHandler(r.logger, r.ResourceService)).Methods(http.MethodDelete).Name(DeleteResourceAPIName)
	s.HandleFunc(GetResourceStatusURI, handlers.UpdateResourceStatusHandler(r.logger, r.ResourceService)).Methods(http.MethodPatch).Name(GetResourceStatusAPIName)
}

func (r *Router) initializeMiddleware(){
	r.Use(middlewares.RecoveryMiddleware(r.logger))
	r.Use(middlewares.TracerMiddleware())
	r.Use(middlewares.InstrumentationMiddleware(r.logger))
}