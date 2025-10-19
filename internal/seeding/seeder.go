package seeding

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/repositories"
	"gorm.io/gorm"
)

func InsertTestData(db *gorm.DB) error {
	// Initialize repositories
	orgRepo := repositories.NewOrganisationRepository(db)
	roleRepo := repositories.NewRoleRepository(db)
	memberRepo := repositories.NewMemberRepository(db)
	teamRepo := repositories.NewTeamRepository(db)
	projectRepo := repositories.NewProjectRepository(db)
	userRoleRepo := repositories.NewUserRoleRepository(db)
	teamMemberRepo := repositories.NewTeamMemberRepository(db)
	memberProjectRepo := repositories.NewMemberProjectRepository(db)

	// Initialize seeders in correct order
	orgSeeder := NewOrganisationSeeder(*orgRepo)
	roleSeeder := NewRoleSeeder(*roleRepo)
	memberSeeder := NewMemberSeeder(*memberRepo, *orgRepo)
	teamSeeder := NewTeamSeeder(*teamRepo, *orgRepo)
	projectSeeder := NewProjectSeeder(*projectRepo, *orgRepo)
	userRoleSeeder := NewUserRoleSeeder(*userRoleRepo, *memberRepo, *roleRepo)
	teamMemberSeeder := NewTeamMemberSeeder(*teamMemberRepo, *teamRepo, *memberRepo)
	memberProjectSeeder := NewMemberProjectSeeder(*memberProjectRepo, *projectRepo, *memberRepo)

	// Create seed manager with correct order (dependencies first)
	seedManager := NewSeedManager(
		orgSeeder,           // 1. Organizations first
		roleSeeder,          // 2. Roles second
		memberSeeder,        // 3. Members (depends on organizations)
		teamSeeder,          // 4. Teams (depends on organizations)
		projectSeeder,       // 5. Projects (depends on organizations)
		userRoleSeeder,      // 6. User roles (depends on members and roles)
		teamMemberSeeder,    // 7. Team members (depends on teams and members)
		memberProjectSeeder, // 8. Member projects (depends on projects and members)
	)

	return seedManager.SeedAll()
}
