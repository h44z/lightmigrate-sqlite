package sqlite

import "fmt"

var (
	// ErrNoDatabaseName signals a missing database name.
	ErrNoDatabaseName = fmt.Errorf("no database name")
	// ErrNoDatabaseClient signals a missing database client.
	ErrNoDatabaseClient = fmt.Errorf("no database client")
)
