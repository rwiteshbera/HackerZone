package database

import _ "database/sql"

// Queries for participants end
var SignUpQuery = "INSERT INTO Participants (uuid, email, first_name, last_name, bio, gender, password, last_login, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
var GetUserAllDataQuery = "SELECT uuid, email, first_name, last_name, bio, gender, password, last_login, created_at FROM Participants WHERE email=$1"
var DeleteParticipantQuery = "DELETE FROM Participants WHERE email=$1"

// Queries for organizers end
var SignupAsOrganizerQuery = "INSERT INTO Organizers (uuid, email, first_name, last_name, bio, gender, password, last_login, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
var GetOrganizersDataQuery = "SELECT uuid, email, first_name, last_name, bio, gender, password, last_login, created_at FROM Organizers WHERE email=$1"
