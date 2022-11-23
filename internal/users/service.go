package users

import (
	"context"

	"github.com/projects/serverless-iam/internal/models"
)

type UsersService struct {
	repo *models.UsersRepo
}

func NewUsersService(repo *models.UsersRepo) *UsersService {
	return &UsersService{repo}
}

func (r UsersService) GetUsers(ctx context.Context) ([]models.User, error) {
	return r.repo.GetUsers(ctx)
}
