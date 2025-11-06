package mongo

import (
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/config"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/entities"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/ports"
)

type RepositoryAdaptor struct {
	logger ports.Logger
	config config.MongoConfig
}
func NewRepositoryAdaptor(logger ports.Logger, config config.MongoConfig) *RepositoryAdaptor {
	return &RepositoryAdaptor{
		logger: logger,
		config: config,
	}
}

func (r *RepositoryAdaptor) SaveResource(resource *entities.Resource) (*entities.Resource, error){
	// Implementation to save resource to MongoDB
	return resource, nil
}	


func (r *RepositoryAdaptor) DeleteResource(id string) error{
	// Implementation to save resource to MongoDB
	return nil
}

func (r *RepositoryAdaptor) GetResource(id string) (*entities.Resource, error){
	// Implementation to save resource to MongoDB
	resource := &entities.Resource{}
	return resource, nil
}

func (r *RepositoryAdaptor) ListResources() ([]entities.Resource, error){
	// Implementation to save resource to MongoDB
	resources := make([]entities.Resource,0)
	return resources, nil
}



func (r *RepositoryAdaptor) UpdateResourceStatus(id string, status string) ( error){
	// Implementation to save resource to MongoDB
	return  nil
}