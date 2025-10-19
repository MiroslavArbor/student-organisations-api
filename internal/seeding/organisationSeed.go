package seeding

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"github.com/MiroslavArbor/student-organisations-api/internal/repositories"
)

type OrganisationSeeder struct {
	repo repositories.OrganisationRepository
}

func NewOrganisationSeeder(repo repositories.OrganisationRepository) *OrganisationSeeder {
	return &OrganisationSeeder{repo: repo}
}

func (s *OrganisationSeeder) Seed() error {
	// Check if organisations already exist
	existingOrgs, err := s.repo.ListAll()
	if err != nil {
		return err
	}
	if len(existingOrgs) > 0 {
		return nil
	}
	orgs := []models.Organisation{
		{Name: "SSFON", Description: "Best organisation in FON"},
		{Name: "Fonis", Description: "Second best organisation in FON"},
		{Name: "Aisec", Description: "Weird people"},
		{Name: "Estiem", Description: "Only woman club"},
		{Name: "SportFON", Description: "Sports organisation"},
	}

	for _, org := range orgs {
		if err := s.repo.Create(&org); err != nil {
			return err
		}
	}
	return nil
}
