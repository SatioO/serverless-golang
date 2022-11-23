package organization

import (
	"context"

	"github.com/projects/serverless-iam/internal/models"
)

type Organization struct {
	repo *models.Organization
}

func NewOrganization(repo *models.Organization) *Organization {
	return &Organization{repo}
}

func (o Organization) CreateOrganization(ctx context.Context, payload *models.CreateOrganization) error {
	return o.repo.CreateOrganization(payload)
}
