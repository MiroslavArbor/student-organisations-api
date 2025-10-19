package seeding

import (
	"time"

	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"github.com/MiroslavArbor/student-organisations-api/internal/repositories"
)

type ProjectSeeder struct {
	projectRepo      repositories.ProjectRepository
	organisationRepo repositories.OrganisationRepository
}

func NewProjectSeeder(projectRepo repositories.ProjectRepository, organisationRepo repositories.OrganisationRepository) *ProjectSeeder {
	return &ProjectSeeder{
		projectRepo:      projectRepo,
		organisationRepo: organisationRepo,
	}
}

func (s *ProjectSeeder) Seed() error {
	// Check if projects already exist
	existingProjects, err := s.projectRepo.ListAll()
	if err != nil {
		return err
	}
	if len(existingProjects) > 0 {
		return nil
	}

	// Get organizations to create projects for
	orgs, err := s.organisationRepo.ListAll()
	if err != nil {
		return err
	}
	if len(orgs) == 0 {
		return nil // No organizations to seed projects for
	}

	// Sample projects for different organizations
	projectData := []struct {
		name        string
		description string
		startDate   time.Time
		endDate     time.Time
		orgIndex    int
	}{
		// SSFON projects
		{"FON Mobile App", "Developing mobile app for FON students", time.Now().AddDate(0, -2, 0), time.Now().AddDate(0, 4, 0), 0},
		{"Tech Conference 2025", "Annual technology conference", time.Now().AddDate(0, 1, 0), time.Now().AddDate(0, 2, 0), 0},
		{"Mentorship Program", "Connecting senior and junior students", time.Now().AddDate(0, 0, -15), time.Now().AddDate(0, 8, 0), 0},

		// Fonis projects
		{"Digital Library", "Online resource library for students", time.Now().AddDate(0, -1, 0), time.Now().AddDate(0, 6, 0), 1},
		{"Career Fair 2025", "Annual career fair event", time.Now().AddDate(0, 2, 0), time.Now().AddDate(0, 3, 0), 1},

		// Aisec projects
		{"Global Exchange Program", "International student exchange", time.Now().AddDate(0, 0, -30), time.Now().AddDate(1, 0, 0), 2},
		{"Leadership Summit", "Leadership development workshop series", time.Now().AddDate(0, 1, 15), time.Now().AddDate(0, 2, 15), 2},

		// Estiem projects
		{"Research Publication", "Academic research on engineering education", time.Now().AddDate(0, -3, 0), time.Now().AddDate(0, 9, 0), 3},
		{"Women in STEM Initiative", "Promoting women participation in STEM", time.Now().AddDate(0, 0, -10), time.Now().AddDate(0, 10, 0), 3},

		// SportFON projects
		{"FON Olympics", "Inter-faculty sports competition", time.Now().AddDate(0, 3, 0), time.Now().AddDate(0, 4, 0), 4},
		{"Fitness Challenge", "Monthly fitness challenges for students", time.Now().AddDate(0, 0, -7), time.Now().AddDate(0, 12, 0), 4},
	}

	for _, projectInfo := range projectData {
		if projectInfo.orgIndex < len(orgs) {
			project := models.Project{
				OrganisationID: orgs[projectInfo.orgIndex].ID,
				Name:           projectInfo.name,
				Description:    projectInfo.description,
				StartDate:      projectInfo.startDate,
				EndDate:        projectInfo.endDate,
			}
			if err := s.projectRepo.Create(&project); err != nil {
				return err
			}
		}
	}
	return nil
}
