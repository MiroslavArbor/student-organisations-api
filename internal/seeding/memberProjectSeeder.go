package seeding

import (
	"time"

	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"github.com/MiroslavArbor/student-organisations-api/internal/repositories"
)

type MemberProjectSeeder struct {
	memberProjectRepo repositories.MemberProjectRepository
	projectRepo       repositories.ProjectRepository
	memberRepo        repositories.MemberRepository
}

func NewMemberProjectSeeder(memberProjectRepo repositories.MemberProjectRepository, projectRepo repositories.ProjectRepository, memberRepo repositories.MemberRepository) *MemberProjectSeeder {
	return &MemberProjectSeeder{
		memberProjectRepo: memberProjectRepo,
		projectRepo:       projectRepo,
		memberRepo:        memberRepo,
	}
}

func (s *MemberProjectSeeder) Seed() error {
	// Check if member projects already exist
	existingMemberProjects, err := s.memberProjectRepo.ListAll()
	if err != nil {
		return err
	}
	if len(existingMemberProjects) > 0 {
		return nil
	}

	// Get projects and members
	projects, err := s.projectRepo.ListAll()
	if err != nil {
		return err
	}
	if len(projects) == 0 {
		return nil // No projects to assign members to
	}

	members, err := s.memberRepo.ListAll()
	if err != nil {
		return err
	}
	if len(members) == 0 {
		return nil // No members to assign to projects
	}

	// Group members by organization for easier assignment
	membersByOrg := make(map[uint][]models.Member)
	for _, member := range members {
		membersByOrg[member.OrganisationID] = append(membersByOrg[member.OrganisationID], member)
	}

	// Available roles for project assignments
	projectRoles := []string{
		"Project Manager",
		"Developer",
		"Designer",
		"Tester",
		"Coordinator",
		"Marketing Lead",
		"Content Creator",
		"Researcher",
		"Team Lead",
		"Contributor",
	}

	// Available grades
	grades := []string{"A", "B", "C", "D", "E", "F", ""}

	// Assign members to projects within their organization
	for i, project := range projects {
		orgMembers, exists := membersByOrg[project.OrganisationID]
		if !exists || len(orgMembers) == 0 {
			continue // No members in this organization
		}

		// Assign 3-5 members to each project (depending on available members)
		membersToAssign := len(orgMembers)
		if membersToAssign > 5 {
			membersToAssign = 5
		}
		if membersToAssign < 3 && len(orgMembers) >= 3 {
			membersToAssign = 3
		}

		for j := 0; j < membersToAssign && j < len(orgMembers); j++ {
			// Assign different roles to members
			role := projectRoles[j%len(projectRoles)]

			// Assign grades randomly (some projects might not have grades yet)
			grade := grades[(i+j)%len(grades)]

			memberProject := models.MemberProject{
				ProjectID:    project.ID,
				MemberID:     orgMembers[j].ID,
				AssignedAt:   time.Now().AddDate(0, 0, -((i * 7) + j)), // Stagger assignment dates
				AssignedRole: role,
				Grade:        grade,
			}
			if err := s.memberProjectRepo.Create(&memberProject); err != nil {
				return err
			}
		}
	}

	return nil
}
