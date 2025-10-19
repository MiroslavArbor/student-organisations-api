package repositories

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"gorm.io/gorm"
)

type TeamRepository struct {
	*BaseRepository[models.Team]
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{
		BaseRepository: NewBaseRepository[models.Team](db),
	}
}
