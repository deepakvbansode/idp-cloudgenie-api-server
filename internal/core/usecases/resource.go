package usecases

import (
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/entities"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/ports"
)

type ResourceService struct {
	logger ports.Logger
	githubProvider ports.GithubPort
	repository ports.RepositoryPort
}

func NewResourceService(logger ports.Logger, githubProvider ports.GithubPort, repository ports.RepositoryPort) *ResourceService {
	return &ResourceService{
		 logger: logger,
		 githubProvider: githubProvider,
		 repository: repository,
	}
}

func (s *ResourceService) CreateResource(resource *entities.Resource) (*entities.Resource, error) {
       // 1. Validate the user (skipped for now)

       // 2. Save the metadata in db (repository)
       savedResource, err := s.repository.SaveResource(resource)
       if err != nil {
	       return nil, err
       }

       // 3. Create XRD from the blueprint selected (placeholder logic)
       xrd := "xrd-content-based-on-blueprint" // TODO: generate XRD from blueprint

       // 4. Push the XRD to github repo
       err = s.githubProvider.PushXRDToRepo(xrd, "repo-name", "path/to/xrd.yaml")
       if err != nil {
	       return nil, err
       }

       return savedResource, nil
}


func (s *ResourceService) UpdateResource(resource *entities.Resource) (*entities.Resource, error) {
	// Save updated resource in db (repository)
	updatedResource, err := s.repository.SaveResource(resource)
	if err != nil {
		 return nil, err
	}

	// Optionally update XRD in github if needed (placeholder logic)
	// xrd := "updated-xrd-content" // TODO: generate updated XRD if required
	// err = s.githubProvider.PushXRDToRepo(xrd, "repo-name", "path/to/xrd.yaml")
	// if err != nil {
	//      return nil, err
	// }

	return updatedResource, nil
}


func (s *ResourceService) DeleteResource(id string) error {
	// Delete resource from db (repository)
	err := s.repository.DeleteResource(id)
	if err != nil {
		 return err
	}

	// Optionally delete XRD from github (not implemented)
	// err = s.githubProvider.DeleteXRDFromRepo("repo-name", "path/to/xrd.yaml")
	// if err != nil {
	//      return err
	// }

	return nil
}


func (s *ResourceService) GetResource(id string) (*entities.Resource, error) {
	// Get resource from db (repository)
	return s.repository.GetResource(id)
}


func (s *ResourceService) ListResources() ([]entities.Resource, error) {
	// List all resources from db (repository)
	return s.repository.ListResources()
}


func (s *ResourceService) UpdateResourceStatus(id string, status string) error {
	// Update status in db (repository)
	return s.repository.UpdateResourceStatus(id, status)
}
