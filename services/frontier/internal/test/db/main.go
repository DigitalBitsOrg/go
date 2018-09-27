// Package db provides helpers to connect to test databases.  It has no
// internal dependencies on frontier and so should be able to be imported by
// any frontier package.
package db

import (
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
	// pq enables postgres support
	db "github.com/digitalbitsorg/go/support/db/dbtest"
	_ "github.com/lib/pq"
)

var (
	coreDB     *sqlx.DB
	coreUrl    *string
	horizonDB  *sqlx.DB
	horizonUrl *string
)

// Horizon returns a connection to the frontier test database
func Horizon(t *testing.T) *sqlx.DB {
	if horizonDB != nil {
		return horizonDB
	}
	postgres := db.Postgres(t)
	horizonUrl = &postgres.DSN
	horizonDB = postgres.Open()
	return horizonDB
}

// HorizonURL returns the database connection the url any test
// use when connecting to the history/frontier database
func HorizonURL() string {
	if horizonUrl == nil {
		log.Panic(fmt.Errorf("Frontier not initialized"))
	}
	return *horizonUrl
}

// StellarCore returns a connection to the digitalbits-core test database
func StellarCore(t *testing.T) *sqlx.DB {
	if coreDB != nil {
		return coreDB
	}
	postgres := db.Postgres(t)
	coreUrl = &postgres.DSN
	coreDB = postgres.Open()
	return coreDB
}

// StellarCoreURL returns the database connection the url any test
// use when connecting to the digitalbits-core database
func StellarCoreURL() string {
	if coreUrl == nil {
		log.Panic(fmt.Errorf("digitalbits-core not initialized"))
	}
	return *coreUrl
}
