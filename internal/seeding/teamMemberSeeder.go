package seeding

import (
	"github.com/MiroslavArbor/student-organisations-api/internal/models"
	"github.com/MiroslavArbor/student-organisations-api/internal/repositories"
)

type TeamMemberSeeder struct {
	teamMemberRepo repositories.TeamMemberRepository
	teamRepo       repositories.TeamRepository
	memberRepo     repositories.MemberRepository
}

func NewTeamMemberSeeder(teamMemberRepo repositories.TeamMemberRepository, teamRepo repositories.TeamRepository, memberRepo repositories.MemberRepository) *TeamMemberSeeder {
	return &TeamMemberSeeder{
		teamMemberRepo: teamMemberRepo,
		teamRepo:       teamRepo,
		memberRepo:     memberRepo,
	}
}

func (s *TeamMemberSeeder) Seed() error {
	// Check if team members already exist
	existingTeamMembers, err := s.teamMemberRepo.ListAll()
	if err != nil {
		return err
	}
	if len(existingTeamMembers) > 0 {
		return nil
	}

	// Get teams and members
	teams, err := s.teamRepo.ListAll()
	if err != nil {
		return err
	}
	if len(teams) == 0 {
		return nil // No teams to assign members to
	}

	members, err := s.memberRepo.ListAll()
	if err != nil {
		return err
	}
	if len(members) == 0 {
		return nil // No members to assign to teams
	}

	// Group members by organization for easier assignment
	membersByOrg := make(map[uint][]models.Member)
	for _, member := range members {
		membersByOrg[member.OrganisationID] = append(membersByOrg[member.OrganisationID], member)
	}

	// Assign members to teams within their organization
	for _, team := range teams {
		orgMembers, exists := membersByOrg[team.OrganisationID]
		if !exists || len(orgMembers) == 0 {
			continue // No members in this organization
		}

		// Assign 2-4 members to each team (depending on available members)
		membersToAssign := len(orgMembers)
		if membersToAssign > 4 {
			membersToAssign = 4
		}
		if membersToAssign < 2 && len(orgMembers) >= 2 {
			membersToAssign = 2
		}

		for i := 0; i < membersToAssign && i < len(orgMembers); i++ {
			teamMember := models.TeamMember{
				TeamID:   team.ID,
				MemberID: orgMembers[i].ID,
			}
			if err := s.teamMemberRepo.Create(&teamMember); err != nil {
				return err
			}
		}
	}

	return nil
}
