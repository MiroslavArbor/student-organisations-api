package seeding

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"github.com/MiroslavArbor/student-organisations-api/internal/repositories"
)

type RoleSeeder struct {
	repo repositories.RoleRepository
}

func NewRoleSeeder(repo repositories.RoleRepository) *RoleSeeder {
	return &RoleSeeder{repo: repo}
}

func (s *RoleSeeder) Seed() error {
	// Check if roles already exist
	existingRoles, err := s.repo.ListAll()
	if err != nil {
		return err
	}
	if len(existingRoles) > 0 {
		return nil
	}

	roles := []models.Role{
		{Name: "Admin", Description: "Full administrative privileges", RoleKey: "admin"},
		{Name: "President", Description: "Organization president", RoleKey: "president"},
		{Name: "Vice President", Description: "Organization vice president", RoleKey: "vice_president"},
		{Name: "Secretary", Description: "Organization secretary", RoleKey: "secretary"},
		{Name: "Treasurer", Description: "Manages organization finances", RoleKey: "treasurer"},
		{Name: "Project Manager", Description: "Manages projects and teams", RoleKey: "project_manager"},
		{Name: "Team Lead", Description: "Leads a specific team", RoleKey: "team_lead"},
		{Name: "Member", Description: "Regular organization member", RoleKey: "member"},
		{Name: "Guest", Description: "Limited access guest user", RoleKey: "guest"},
	}

	for _, role := range roles {
		if err := s.repo.Create(&role); err != nil {
			return err
		}
	}
	return nil
}
