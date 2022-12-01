package realm

import "github.com/projects/serverless-iam/internal/models"

type realmService struct {
	repo *models.RealmRepo
}

func NewRealmService(repo *models.RealmRepo) *realmService {
	return &realmService{repo}
}

func (r realmService) GetRealms() ([]models.Realm, error) {
	return r.repo.GetRealms()
}
