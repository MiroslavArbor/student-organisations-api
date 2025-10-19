package repositories

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"gorm.io/gorm"
)

type TeamMemberRepository struct {
	*BaseRepository[models.TeamMember]
}

func NewTeamMemberRepository(db *gorm.DB) *TeamMemberRepository {
	return &TeamMemberRepository{
		BaseRepository: NewBaseRepository[models.TeamMember](db),
	}
}

func (r *TeamMemberRepository) FindByTeamID(teamID uint) ([]models.TeamMember, error) {
	var teamMembers []models.TeamMember
	if err := r.db.Where("team_id = ?", teamID).Find(&teamMembers).Error; err != nil {
		return nil, err
	}
	return teamMembers, nil
}

func (r *TeamMemberRepository) FindByMemberID(memberID uint) ([]models.TeamMember, error) {
	var teamMembers []models.TeamMember
	if err := r.db.Where("member_id = ?", memberID).Find(&teamMembers).Error; err != nil {
		return nil, err
	}
	return teamMembers, nil
}

func (r *TeamMemberRepository) AddMemberToTeam(teamID, memberID uint) error {
	teamMember := models.TeamMember{
		TeamID:   teamID,
		MemberID: memberID,
	}
	return r.db.Create(&teamMember).Error
}

func (r *TeamMemberRepository) RemoveMemberFromTeam(teamID, memberID uint) error {
	return r.db.Where("team_id = ? AND member_id = ?", teamID, memberID).Delete(&models.TeamMember{}).Error
}

func (r *TeamMemberRepository) IsMemberInTeam(teamID, memberID uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.TeamMember{}).Where("team_id = ? AND member_id = ?", teamID, memberID).Count(&count).Error
	return count > 0, err
}
