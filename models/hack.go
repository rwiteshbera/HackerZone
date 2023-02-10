package models

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
