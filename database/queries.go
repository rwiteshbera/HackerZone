package database

import _ "database/sql"

// Queries for participants end
var SignUpQuery = "INSERT INTO participants (email, firstName, lastName, password, lastLogin, createdAt) VALUES ($1, $2, $3, $4, $5, $6)"
var GetUserAllDataQuery = "SELECT email, firstName, lastName, password, lastLogin, createdAt FROM participants WHERE email=$1"
var DeleteParticipantQuery = "DELETE FROM participants WHERE email=$1"

// Queries for organizers end
var SignupAsOrganizerQuery = "INSERT INTO organizers (email, firstName, lastName, password, lastLogin, createdAt) VALUES ($1, $2, $3, $4, $5, $6)"
var GetOrganizersDataQuery = "SELECT email, firstName, lastName, password, lastLogin, createdAt FROM organizers WHERE email=$1"
