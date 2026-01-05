package gobizfly

import "context"

const (
	projectPath = "/projects"
)

type iamService struct {
	client *Client
}

var _ IAMService = (*iamService)(nil)

type IAMService interface {
	ListProjects(ctx context.Context, opts ListProjectsOpts) ([]*IAMProject, error)
}

func (i iamService) projectResourcePath() string {
	return projectPath
}
