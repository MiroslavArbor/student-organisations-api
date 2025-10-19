package repositories

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"gorm.io/gorm"
)

type MemberProjectRepository struct {
	*BaseRepository[models.MemberProject]
}

func NewMemberProjectRepository(db *gorm.DB) *MemberProjectRepository {
	return &MemberProjectRepository{
		BaseRepository: NewBaseRepository[models.MemberProject](db),
	}
}

func (r *MemberProjectRepository) FindByProjectID(projectID uint) ([]models.MemberProject, error) {
	var memberProjects []models.MemberProject
	if err := r.db.Where("project_id = ?", projectID).Find(&memberProjects).Error; err != nil {
		return nil, err
	}
	return memberProjects, nil
}

func (r *MemberProjectRepository) FindByMemberID(memberID uint) ([]models.MemberProject, error) {
	var memberProjects []models.MemberProject
	if err := r.db.Where("member_id = ?", memberID).Find(&memberProjects).Error; err != nil {
		return nil, err
	}
	return memberProjects, nil
}

func (r *MemberProjectRepository) AssignMemberToProject(memberProject *models.MemberProject) error {
	return r.db.Create(memberProject).Error
}

func (r *MemberProjectRepository) RemoveMemberFromProject(projectID, memberID uint) error {
	return r.db.Where("project_id = ? AND member_id = ?", projectID, memberID).Delete(&models.MemberProject{}).Error
}

func (r *MemberProjectRepository) UpdateAssignment(memberProject *models.MemberProject) error {
	return r.db.Model(memberProject).Where("project_id = ? AND member_id = ?", memberProject.ProjectID, memberProject.MemberID).Updates(memberProject).Error
}

func (r *MemberProjectRepository) IsMemberAssigned(projectID, memberID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.MemberProject{}).Where("project_id = ? AND member_id = ?", projectID, memberID).Count(&count).Error
	return count > 0, err
}

func (r *MemberProjectRepository) FindByRole(role string) ([]models.MemberProject, error) {
	var memberProjects []models.MemberProject
	if err := r.db.Where("assigned_role = ?", role).Find(&memberProjects).Error; err != nil {
		return nil, err
	}
	return memberProjects, nil
}

func (r *MemberProjectRepository) FindByGrade(grade string) ([]models.MemberProject, error) {
	var memberProjects []models.MemberProject
	if err := r.db.Where("grade = ?", grade).Find(&memberProjects).Error; err != nil {
		return nil, err
	}
	return memberProjects, nil
}
