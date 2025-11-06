
package ports

import "github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/entities"

// RepositoryPort defines the interface for DB operations for resources
type RepositoryPort interface {
	SaveResource(resource *entities.Resource) (*entities.Resource, error)
	DeleteResource(id string) error
	GetResource(id string) (*entities.Resource, error)
	ListResources() ([]entities.Resource, error)
	UpdateResourceStatus(id string, status string) ( error)
}

// GithubPort defines the interface for Github operations
type GithubPort interface {
	PushXRDToRepo(xrd string, repo string, path string) error
	// ...other github methods as needed
}
