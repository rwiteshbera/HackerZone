package database

import _ "database/sql"

var SignUpQuery = "INSERT INTO users (email, firstName, lastName, password, lastLogin, createdAt) VALUES ($1, $2, $3, $4, $5, $6)"
var GetUserAllDataQuery = "SELECT email, firstName, lastName, password, lastLogin, createdAt FROM users WHERE email=$1"
