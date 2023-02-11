package database

import _ "database/sql"

// Queries for participants end
var SignUpQuery = `INSERT INTO Participants (uuid, email, first_name, last_name, bio, gender, password, last_login, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
var GetUserAllDataByEmailQuery = `SELECT uuid, email, first_name, last_name, bio, gender, password, last_login, created_at FROM Participants WHERE email=$1`
var DeleteParticipantQuery = `DELETE FROM Participants WHERE email=$1`

// Queries for organizers end
var SignupAsOrganizerQuery = `INSERT INTO Organizers (uuid, email, first_name, last_name, bio, gender, password, last_login, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
var GetOrganizersDataByEmailQuery = `SELECT uuid, email, first_name, last_name, bio, gender, password, last_login, created_at FROM Organizers WHERE email=$1`

// Check hackathon name availability
var HackathonNameCheckQuery = `SELECT id FROM Hackathons WHERE name=$1`
var CreateHackathonQuery = `INSERT INTO Hackathons (name, tagline, description, contact_email, host, hacking_start, deadline,
created_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

// Create Team
var CreateTeamQuery = `INSERT INTO Teams (team_id, hackathon_id, name, created_by) VALUES ($1, $2, $3, $4)`

// Get All Team Details You are a member
var GetAllYourTeamDetailsQuery = ``

// Add Member in Team
var InsertMemberInTeamQuery = `INSERT INTO Team_Members (team_id, email, hackathon_id) VALUES($1, $2, $3)`

// Delete Member from Team
var DeleteMemberQuery = `DELETE FROM Team_Members WHERE team_id = $1 AND email = $2`
