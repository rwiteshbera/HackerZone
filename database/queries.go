package database

import _ "database/sql"

var SignUpQuery = "INSERT INTO participants (email, firstName, lastName, password, lastLogin, createdAt) VALUES ($1, $2, $3, $4, $5, $6)"
var GetUserAllDataQuery = "SELECT email, firstName, lastName, password, lastLogin, createdAt FROM participants WHERE email=$1"
var DeleteParticipantQuery = "DELETE FROM participants WHERE email=$1"
