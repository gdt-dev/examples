package db_test

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	"github.com/gdt-dev/gdt"
	gdtfix "github.com/gdt-dev/gdt/fixture"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

const (
	testDBPath = "./test.db"
	initDB     = `
PRAGMA synchronous=FULL;
PRAGMA foreign_keys=1;

CREATE TABLE users (
  id INTEGER PRIMARY KEY NOT NULL
, name TEXT NOT NULL
, created_on TEXT NOT NULL
);
CREATE INDEX ix_users_created ON users(created_on);

INSERT INTO users (id, name, created_on) VALUES (1, 'Alice', DATETIME('2009-01-03'));
INSERT INTO users (id, name, created_on) VALUES (2, 'Bob', DATETIME('2015-11-16'));
INSERT INTO users (id, name, created_on) VALUES (3, 'Chuck', DATETIME('2019-08-07'));
`
)

// resetDB clears any data in our database and populates the database with our
// base starting data set
func resetDB() {
	os.Remove(testDBPath)
	db, err := sql.Open("sqlite3", testDBPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec(initDB)
	if err != nil {
		panic(err)
	}
}

func TestDataAccessLayer(t *testing.T) {
	require := require.New(t)

	fp := filepath.Join("testdata", "data-access.yaml")

	s, err := gdt.From(fp)
	require.Nil(err)

	dbFix := gdtfix.New(gdtfix.WithStarter(resetDB))

	ctx := gdt.NewContext()
	ctx = gdt.RegisterFixture(ctx, "db", dbFix)

	err = s.Run(ctx, t)
	require.Nil(err)
}
