package database

import (
	"database/sql"
	"fmt"
	"os"

  _ "github.com/libsql/libsql-client-go/libsql"
	_ "modernc.org/sqlite"
)

var db = "closing-becatron-bahaha"
var dbUrl = fmt.Sprintf("libsql://%s.turso.io?authToken=%s", db, os.Getenv("TURSO_AUTH"))

func DataSource() (*sql.DB, error) {
  return sql.Open("libsql", dbUrl)
}

