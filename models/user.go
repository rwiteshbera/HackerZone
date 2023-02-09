package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UUID      uuid.UUID `json:"uuid"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Bio       string    `json:"bio"`
	Gender    string    `json:"gender"`
	Password  string    `json:"password"`
	LastLogin time.Time `json:"lastLogin"`
	CreatedAt time.Time `json:"createdAt"`
}

// User Signup Request
type SignupRequest struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}

// User Login Request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// User Login Response
type LoginResponse struct {
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Bio       string    `json:"bio"`
	Gender    string    `json:"gender"`
	LastLogin time.Time `json:"lastLogin"`
	CreatedAt time.Time `json:"createdAt"`
}

// Organizer
