package repositories

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"gorm.io/gorm"
)

type ProjectRepository struct {
	*BaseRepository[models.Project]
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		BaseRepository: NewBaseRepository[models.Project](db),
	}
}

func (r *ProjectRepository) FindByOrganisationID(organisationID uint) ([]models.Project, error) {
	var projects []models.Project
	if err := r.db.Where("organisation_id = ?", organisationID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (r *ProjectRepository) FindByName(name string) (*models.Project, error) {
	var project models.Project
	if err := r.db.Where("name = ?", name).First(&project).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepository) FindWithAssignments(id uint) (*models.Project, error) {
	var project models.Project
	if err := r.db.Preload("Assignments").First(&project, id).Error; err != nil {
		return nil, err
	}
	return &project, nil
}

func (r *ProjectRepository) FindActiveProjects() ([]models.Project, error) {
	var projects []models.Project
	if err := r.db.Where("end_date >= CURRENT_DATE").Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}
