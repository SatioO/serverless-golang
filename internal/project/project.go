package project

import (
	"context"

	"github.com/projects/serverless-iam/internal/models"
)

type project struct {
	repo *models.ProjectRepo
}

func NewProject(repo *models.ProjectRepo) *project {
	return &project{repo}
}

func (p project) CreateProject(ctx context.Context, orgId string, payload *models.CreateProject) error {
	return p.repo.CreateProject(orgId, payload)
}
