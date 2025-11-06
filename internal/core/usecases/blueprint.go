package usecases

import (
	"context"

	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/common/constants"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/entities"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/ports"
)

type BlueprintService struct {
	logger ports.Logger
	crossplane ports.CrossplanePort
}

func NewBlueprintService(logger ports.Logger,crossplane ports.CrossplanePort) *BlueprintService {
	return &BlueprintService{
		logger: logger,
		crossplane: crossplane,
	}
}

func (s *BlueprintService) GetBlueprints(ctx context.Context) ([]entities.Blueprint, error) {
	log := s.logger.WithField("tradeId", ctx.Value(constants.TraceIDKey))
	log.Info("Fetching blueprints from Crossplane")
	return s.crossplane.ListBlueprints(ctx)
}
