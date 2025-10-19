package seeding

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"github.com/MiroslavArbor/student-organisations-api/internal/repositories"
)

type TeamSeeder struct {
	teamRepo         repositories.TeamRepository
	organisationRepo repositories.OrganisationRepository
}

func NewTeamSeeder(teamRepo repositories.TeamRepository, organisationRepo repositories.OrganisationRepository) *TeamSeeder {
	return &TeamSeeder{
		teamRepo:         teamRepo,
		organisationRepo: organisationRepo,
	}
}

func (s *TeamSeeder) Seed() error {
	// Check if teams already exist
	existingTeams, err := s.teamRepo.ListAll()
	if err != nil {
		return err
	}
	if len(existingTeams) > 0 {
		return nil
	}

	// Get organizations to create teams for
	orgs, err := s.organisationRepo.ListAll()
	if err != nil {
		return err
	}
	if len(orgs) == 0 {
		return nil // No organizations to seed teams for
	}

	// Sample teams for different organizations
	teamData := []struct {
		name        string
		description string
		orgIndex    int
	}{
		// SSFON teams
		{"Development Team", "Software development and programming", 0},
		{"Marketing Team", "Social media and event promotion", 0},
		{"Events Team", "Organizing workshops and conferences", 0},

		// Fonis teams
		{"Technical Team", "Technical support and infrastructure", 1},
		{"Content Team", "Creating educational content", 1},

		// Aisec teams
		{"International Team", "Managing international partnerships", 2},
		{"Recruitment Team", "Member recruitment and onboarding", 2},

		// Estiem teams
		{"Research Team", "Academic research and publications", 3},
		{"Networking Team", "Building professional networks", 3},

		// SportFON teams
		{"Football Team", "Organizing football tournaments", 4},
		{"Basketball Team", "Managing basketball events", 4},
		{"Fitness Team", "Fitness programs and workshops", 4},
	}

	for _, teamInfo := range teamData {
		if teamInfo.orgIndex < len(orgs) {
			team := models.Team{
				OrganisationID: orgs[teamInfo.orgIndex].ID,
				Name:           teamInfo.name,
				Description:    teamInfo.description,
			}
			if err := s.teamRepo.Create(&team); err != nil {
				return err
			}
		}
	}
	return nil
}
