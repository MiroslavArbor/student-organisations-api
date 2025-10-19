package models

import "time"

type Organisation struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"size:255;not null;unique" json:"name"`
	Description string    `gorm:"size:500" json:"description"`
	Teams       []Team    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"teams"`
	Members     []Member  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members"`
	Projects    []Project `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"projects"`
}

type Role struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"size:100;not null;unique" json:"name"`
	Description string     `gorm:"size:255" json:"description"`
	RoleKey     string     `gorm:"size:100;not null;unique" json:"role_key"`
	Users       []UserRole `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Member struct {
	ID             uint            `gorm:"primaryKey" json:"id"`
	OrganisationID uint            `gorm:"not null" json:"organisation_id"`
	Name           string          `gorm:"size:255;not null" json:"name"`
	Email          string          `gorm:"size:255;not null;unique" json:"email"`
	PasswordHash   string          `gorm:"size:255;not null" json:"-"`
	Roles          []UserRole      `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"roles"`
	Teams          []TeamMember    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"teams"`
	Projects       []MemberProject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"projects"`
}

type UserRole struct {
	MemberID uint `gorm:"primaryKey" json:"member_id"`
	RoleID   uint `gorm:"primaryKey" json:"role_id"`
	Role     Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type Team struct {
	ID             uint         `gorm:"primaryKey" json:"id"`
	OrganisationID uint         `gorm:"not null" json:"organisation_id"`
	Name           string       `gorm:"size:255;not null" json:"name"`
	Description    string       `gorm:"size:500" json:"description"`
	Members        []TeamMember `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"members"`
}

type TeamMember struct {
	TeamID   uint `gorm:"primaryKey" json:"team_id"`
	MemberID uint `gorm:"primaryKey" json:"member_id"`
}

type Project struct {
	ID             uint            `gorm:"primaryKey" json:"id"`
	OrganisationID uint            `gorm:"not null" json:"organisation_id"`
	Name           string          `gorm:"size:255;not null" json:"name"`
	Description    string          `gorm:"size:500" json:"description"`
	StartDate      time.Time       `json:"start_date"`
	EndDate        time.Time       `json:"end_date"`
	Assignments    []MemberProject `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"assignments"`
}

type MemberProject struct {
	ProjectID    uint      `gorm:"primaryKey" json:"project_id"`
	MemberID     uint      `gorm:"primaryKey" json:"member_id"`
	AssignedAt   time.Time `json:"assigned_at"`
	AssignedRole string    `gorm:"size:100" json:"assigned_role"`
	Grade        string    `gorm:"size:50" json:"grade"`
}
