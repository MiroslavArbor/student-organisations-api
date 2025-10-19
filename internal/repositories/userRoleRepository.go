package repositories

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"gorm.io/gorm"
)

type UserRoleRepository struct {
	*BaseRepository[models.UserRole]
}

func NewUserRoleRepository(db *gorm.DB) *UserRoleRepository {
	return &UserRoleRepository{
		BaseRepository: NewBaseRepository[models.UserRole](db),
	}
}

func (r *UserRoleRepository) FindByMemberID(memberID uint) ([]models.UserRole, error) {
	var userRoles []models.UserRole
	if err := r.db.Where("member_id = ?", memberID).Preload("Role").Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (r *UserRoleRepository) FindByRoleID(roleID uint) ([]models.UserRole, error) {
	var userRoles []models.UserRole
	if err := r.db.Where("role_id = ?", roleID).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	return userRoles, nil
}

func (r *UserRoleRepository) AssignRole(memberID, roleID uint) error {
	userRole := models.UserRole{
		MemberID: memberID,
		RoleID:   roleID,
	}
	return r.db.Create(&userRole).Error
}

func (r *UserRoleRepository) RemoveRole(memberID, roleID uint) error {
	return r.db.Where("member_id = ? AND role_id = ?", memberID, roleID).Delete(&models.UserRole{}).Error
}

func (r *UserRoleRepository) HasRole(memberID, roleID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.UserRole{}).Where("member_id = ? AND role_id = ?", memberID, roleID).Count(&count).Error
	return count > 0, err
}
