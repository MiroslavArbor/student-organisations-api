package repositories

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"gorm.io/gorm"
)

type OrganisationRepository struct {
	*BaseRepository[models.Organisation]
}

func NewOrganisationRepository(db *gorm.DB) *OrganisationRepository {
	return &OrganisationRepository{
		BaseRepository: NewBaseRepository[models.Organisation](db),
	}
}
func (r *OrganisationRepository) FindByName(name string) (*models.Organisation, error) {
	var org models.Organisation
	if err := r.db.Where("name = ?", name).First(&org).Error; err != nil {
		return nil, err
	}
	return &org, nil
}
