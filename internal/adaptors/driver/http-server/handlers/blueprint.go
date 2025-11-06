package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/common/constants"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/ports"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/usecases"
)

func GetBlueprintsHandler(logger ports.Logger, blueprintService *usecases.BlueprintService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx:= r.Context()
		log := logger.WithField("tradeId", ctx.Value(constants.TraceIDKey))
		log.Info("GetBlueprintsHandler called")
		blueprints, err := blueprintService.GetBlueprints(ctx)
		if err != nil {
			log.Error("Failed to get blueprints: ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(blueprints)
	}
}
