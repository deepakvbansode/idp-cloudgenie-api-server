package handlers

import (
	"net/http"

	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/ports"
)

func HealthCheckHandler(logger ports.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, request *http.Request) {
		//code to handle
	}
}
