package database

import _ "database/sql"

var SignUpQuery = "INSERT INTO users (email, firstName, lastName, password, createdAt, lastLogin) VALUES ($1, $2, $3, $4, $5, $6)"
var CheckIfUserAlreadyExistsQuery = "SELECT email, password FROM users WHERE email= $1"
