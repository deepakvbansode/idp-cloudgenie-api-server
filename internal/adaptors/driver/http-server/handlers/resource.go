package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/ports"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/usecases"
)

func GetResourcesHandler(logger ports.Logger, resourceService *usecases.ResourceService) http.HandlerFunc {
       return func(w http.ResponseWriter, r *http.Request) {
	       resources, err := resourceService.ListResources()
	       if err != nil {
		       w.WriteHeader(http.StatusInternalServerError)
		       return
	       }
	       w.Header().Set("Content-Type", "application/json")
	       w.WriteHeader(http.StatusOK)
	       json.NewEncoder(w).Encode(resources)
       }
}

func GetResourceHandler(logger ports.Logger, resourceService *usecases.ResourceService) http.HandlerFunc {
       return func(w http.ResponseWriter, r *http.Request) {
	       // TODO: extract id from URL path
		   logger.Panic("GetResourceHandler not implemented yet")
	       w.WriteHeader(http.StatusNotImplemented)
       }
}

func CreateResourceHandler(logger ports.Logger, resourceService *usecases.ResourceService) http.HandlerFunc {
       return func(w http.ResponseWriter, r *http.Request) {
	       // TODO: parse request body and call CreateResource
	       w.WriteHeader(http.StatusNotImplemented)
       }
}

func DeleteResourceHandler(logger ports.Logger, resourceService *usecases.ResourceService) http.HandlerFunc {
       return func(w http.ResponseWriter, r *http.Request) {
	       // TODO: extract id from URL path and call DeleteResource
	       w.WriteHeader(http.StatusNotImplemented)
       }
}

func UpdateResourceStatusHandler(logger ports.Logger,resourceService *usecases.ResourceService) http.HandlerFunc {
       return func(w http.ResponseWriter, r *http.Request) {
	       // TODO: extract id and status from URL/path/body and call UpdateResourceStatus
	       w.WriteHeader(http.StatusNotImplemented)
       }
}
