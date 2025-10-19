package repositories

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	*BaseRepository[models.Role]
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		BaseRepository: NewBaseRepository[models.Role](db),
	}
}

func (r *RoleRepository) FindByName(name string) (*models.Role, error) {
	var role models.Role
	if err := r.db.Where("name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *RoleRepository) FindWithUsers(id uint) (*models.Role, error) {
	var role models.Role
	if err := r.db.Preload("Users").First(&role, id).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
