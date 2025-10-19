package seeding

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"github.com/MiroslavArbor/student-organisations-api/internal/repositories"
)

type UserRoleSeeder struct {
	userRoleRepo repositories.UserRoleRepository
	memberRepo   repositories.MemberRepository
	roleRepo     repositories.RoleRepository
}

func NewUserRoleSeeder(userRoleRepo repositories.UserRoleRepository, memberRepo repositories.MemberRepository, roleRepo repositories.RoleRepository) *UserRoleSeeder {
	return &UserRoleSeeder{
		userRoleRepo: userRoleRepo,
		memberRepo:   memberRepo,
		roleRepo:     roleRepo,
	}
}

func (s *UserRoleSeeder) Seed() error {
	// Check if user roles already exist
	existingUserRoles, err := s.userRoleRepo.ListAll()
	if err != nil {
		return err
	}
	if len(existingUserRoles) > 0 {
		return nil
	}

	// Get members and roles
	members, err := s.memberRepo.ListAll()
	if err != nil {
		return err
	}
	if len(members) == 0 {
		return nil // No members to assign roles to
	}

	roles, err := s.roleRepo.ListAll()
	if err != nil {
		return err
	}
	if len(roles) == 0 {
		return nil // No roles to assign
	}

	// Find specific roles by name for easier assignment
	roleMap := make(map[string]uint)
	for _, role := range roles {
		roleMap[role.Name] = role.ID
	}

	// Sample role assignments - assign roles to specific members
	roleAssignments := []struct {
		memberIndex int
		roleName    string
	}{
		// SSFON members (first 4 members)
		{0, "President"},       // Marko Petrovic
		{1, "Vice President"},  // Ana Nikolic
		{2, "Project Manager"}, // Stefan Jovanovic
		{3, "Member"},          // Milica Stojanovic

		// Fonis members (next 3 members)
		{4, "President"}, // Petar Milic
		{5, "Secretary"}, // Jovana Radovic
		{6, "Member"},    // Nemanja Popovic

		// Aisec members (next 3 members)
		{7, "President"}, // Marija Antic
		{8, "Treasurer"}, // Aleksandar Savic
		{9, "Member"},    // Teodora Milosevic

		// Estiem members (next 3 members)
		{10, "President"}, // Jelena Radic
		{11, "Team Lead"}, // Katarina Stankovic
		{12, "Member"},    // Tamara Djordjevic

		// SportFON members (next 3 members)
		{13, "President"}, // Milos Nikolic
		{14, "Team Lead"}, // Dusan Markovic
		{15, "Member"},    // Luka Petrovic
	}

	// Assign roles to members
	for _, assignment := range roleAssignments {
		if assignment.memberIndex < len(members) {
			roleID, exists := roleMap[assignment.roleName]
			if exists {
				userRole := models.UserRole{
					MemberID: members[assignment.memberIndex].ID,
					RoleID:   roleID,
				}
				if err := s.userRoleRepo.Create(&userRole); err != nil {
					return err
				}
			}
		}
	}

	// Assign "Member" role to all members who don't have any role yet
	memberRoleID, exists := roleMap["Member"]
	if exists {
		for i, member := range members {
			// Skip members who already have roles assigned
			skipMember := false
			for _, assignment := range roleAssignments {
				if assignment.memberIndex == i {
					skipMember = true
					break
				}
			}
			if !skipMember {
				userRole := models.UserRole{
					MemberID: member.ID,
					RoleID:   memberRoleID,
				}
				if err := s.userRoleRepo.Create(&userRole); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
