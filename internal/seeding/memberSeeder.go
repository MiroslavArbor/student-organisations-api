package seeding

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"github.com/MiroslavArbor/student-organisations-api/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type MemberSeeder struct {
	memberRepo       repositories.MemberRepository
	organisationRepo repositories.OrganisationRepository
}

func NewMemberSeeder(memberRepo repositories.MemberRepository, organisationRepo repositories.OrganisationRepository) *MemberSeeder {
	return &MemberSeeder{
		memberRepo:       memberRepo,
		organisationRepo: organisationRepo,
	}
}

func (s *MemberSeeder) Seed() error {
	// Check if members already exist
	existingMembers, err := s.memberRepo.ListAll()
	if err != nil {
		return err
	}
	if len(existingMembers) > 0 {
		return nil
	}

	// Get organizations to assign members to
	orgs, err := s.organisationRepo.ListAll()
	if err != nil {
		return err
	}
	if len(orgs) == 0 {
		return nil // No organizations to seed members for
	}

	// Hash password for all members
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Sample members for different organizations
	memberData := []struct {
		name     string
		email    string
		orgIndex int
	}{
		// SSFON members
		{"Marko Petrovic", "marko.petrovic@ssfon.rs", 0},
		{"Ana Nikolic", "ana.nikolic@ssfon.rs", 0},
		{"Stefan Jovanovic", "stefan.jovanovic@ssfon.rs", 0},
		{"Milica Stojanovic", "milica.stojanovic@ssfon.rs", 0},

		// Fonis members
		{"Petar Milic", "petar.milic@fonis.rs", 1},
		{"Jovana Radovic", "jovana.radovic@fonis.rs", 1},
		{"Nemanja Popovic", "nemanja.popovic@fonis.rs", 1},

		// Aisec members
		{"Marija Antic", "marija.antic@aisec.rs", 2},
		{"Aleksandar Savic", "aleksandar.savic@aisec.rs", 2},
		{"Teodora Milosevic", "teodora.milosevic@aisec.rs", 2},

		// Estiem members
		{"Jelena Radic", "jelena.radic@estiem.rs", 3},
		{"Katarina Stankovic", "katarina.stankovic@estiem.rs", 3},
		{"Tamara Djordjevic", "tamara.djordjevic@estiem.rs", 3},

		// SportFON members
		{"Milos Nikolic", "milos.nikolic@sportfon.rs", 4},
		{"Dusan Markovic", "dusan.markovic@sportfon.rs", 4},
		{"Luka Petrovic", "luka.petrovic@sportfon.rs", 4},
	}

	for _, memberInfo := range memberData {
		if memberInfo.orgIndex < len(orgs) {
			member := models.Member{
				OrganisationID: orgs[memberInfo.orgIndex].ID,
				Name:           memberInfo.name,
				Email:          memberInfo.email,
				PasswordHash:   string(hashedPassword),
			}
			if err := s.memberRepo.Create(&member); err != nil {
				return err
			}
		}
	}
	return nil
}
