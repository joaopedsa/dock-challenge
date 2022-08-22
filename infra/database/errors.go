package database

import (
	"errors"
)

var (
	ErrDBNil = errors.New("the database is not instantiated")
)

// Error for logging
var (
	StrGetDBFail = "database: failed to get a new database connection"
)
