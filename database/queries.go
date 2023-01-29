package database

import _ "database/sql"

var SignUpQuery = "INSERT INTO users (email, firstName, lastName, password, createdAt, lastLogin) VALUES ($1, $2, $3, $4, $5, $6)"
