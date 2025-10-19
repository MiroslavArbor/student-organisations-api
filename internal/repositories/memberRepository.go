package repositories

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"gorm.io/gorm"
)

type MemberRepository struct {
	*BaseRepository[models.Member]
}

func NewMemberRepository(db *gorm.DB) *MemberRepository {
	return &MemberRepository{
		BaseRepository: NewBaseRepository[models.Member](db),
	}
}

func (r *MemberRepository) FindByEmail(email string) (*models.Member, error) {
	var member models.Member
	if err := r.db.Where("email = ?", email).First(&member).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *MemberRepository) FindByOrganisationID(organisationID uint) ([]models.Member, error) {
	var members []models.Member
	if err := r.db.Where("organisation_id = ?", organisationID).Find(&members).Error; err != nil {
		return nil, err
	}
	return members, nil
}

func (r *MemberRepository) FindWithRoles(id uint) (*models.Member, error) {
	var member models.Member
	if err := r.db.Preload("Roles.Role").First(&member, id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *MemberRepository) FindWithTeams(id uint) (*models.Member, error) {
	var member models.Member
	if err := r.db.Preload("Teams").First(&member, id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *MemberRepository) FindWithProjects(id uint) (*models.Member, error) {
	var member models.Member
	if err := r.db.Preload("Projects").First(&member, id).Error; err != nil {
		return nil, err
	}
	return &member, nil
}
