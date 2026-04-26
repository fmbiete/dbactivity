package database

import (
	"fmt"
	"strings"
)

type DatabaseType int

const (
	Unknown DatabaseType = iota
	Oracle
	MySQL
	PostgreSQL
)

// String defines how the value is displayed (required for flag.Value)
func (d *DatabaseType) String() string {
	switch *d {
	case Oracle:
		return "oracle"
	case MySQL:
		return "mysql"
	case PostgreSQL:
		return "postgresql"
	default:
		return "unknown"
	}
}

// Set maps the string flag input to your integer constants
func (d *DatabaseType) Set(value string) error {
	switch strings.ToLower(value) {
	case "oracle":
		*d = Oracle
	case "mysql":
		*d = MySQL
	case "postgresql":
		*d = PostgreSQL
	default:
		return fmt.Errorf("must be oracle, mysql, or postgresql")
	}
	return nil
}
