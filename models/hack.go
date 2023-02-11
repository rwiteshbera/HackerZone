package models

import "github.com/google/uuid"

type Hackathon struct {
	Id           int8   `json:"id"`
	Name         string `json:"name"`
	Tagline      string `json:"tagline"`
	Description  string `json:"description"`
	ContactEmail string `json:"contactEmail"`
	Host         string `json:"host"`
	HackingStart string `json:"hackingStart"`
	Deadline     string `json:"deadline"`
	CreatedBy    string `json:"createdBy"`
}

type Team struct {
	TeamId      uuid.UUID `json:"teamId"`
	HackathonId int8      `json:"hackathonId"`
	Name        string    `json:"name"`
	CreatedBy   uuid.UUID `json:"createdBy"`
}

type TeamMember struct {
	TeamId      uuid.UUID `json:"teamId"`
	MemberEmail string    `json:"memberEmail"`
	HackathonId int8      `json:"hackathonId"`
}
