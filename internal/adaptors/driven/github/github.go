package github

import (
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/config"
	"github.com/deepakvbansode/idp-cloudgenie-backend/internal/core/ports"
)

type GithubAdaptor struct {
	logger ports.Logger
	config config.GithubConfig
}

func NewGithubAdaptor(logger ports.Logger,config config.GithubConfig) *GithubAdaptor {
	return &GithubAdaptor{
		logger: logger,
		config: config,
	}
}

func (g *GithubAdaptor) PushXRDToRepo(xrd string, repo string, path string) error {
	// Implementation to push XRD to GitHub repository
	return nil
}	