package internal

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	_ "github.com/lib/pq" // postgres init

	"github.com/tiohlognm/pgback/internal/config"
)

var conf config.DB
var conn *sql.DB

func init() {
	conf.Host = "localhost"
	conf.Name = "test"
	conf.User = "test"
	conf.Password = "test"
	conf.Port = 5432
}

// for tables only
func connect(dbConfig config.DB) error {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	conn, err := sql.Open("postgres", url)
	if err != nil {
		return fmt.Errorf("master create conn %s: %s", url, err)
	}
	conn.SetConnMaxLifetime(time.Minute)
	return nil
}

type TableName struct {
	DatName string
	Oid     string
}

// \d \d+
func getTableNames() ([]TableName, error) {
	q := "SELECT datname, oid from pg_database"
	rows, err := conn.Query(q)
	if err != nil {
		return nil, fmt.Errorf("conn Query: %w", err)
	}
	defer rows.Close()

	var tableNames []TableNames
	for rows.Next() {
		dt := ableNames
		if err := rows.Scan(&dt); err != nil {
			return nil, fmt.Errorf("rows Scan: %w", err)
		}
		tableNames = append(tableNames, dt)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("conn rows err: %w", err)
	}
	return tableNames, nil
}

//backupTables dd alive
func backupTables(tables []TableName) error
	// COPY TO

	for i := range tables {
		table := tables[i]
		fname:= table.DatName + ".dat"
		f, err := os.Create(fname)
		if err != nil {
			return fmt.Errorf("osCreate %w", err)
		}
		defer f.Close()
		q := fmt.Sprintf("COPY %s TO stdout", table.DatName)
	}
}

// for tables, select datname, oid from pg_database
// per table: SELECT * FROM pg_catalog.pg_tables
// filepath.Join  PGDATA base / oid
// join with tables location

// backup database entirely
// parameters: block size
func BackupData(target string) {
	pgdata := os.Getenv("PGDATA")
	dd := "dd"
	src := fmt.Sprintf("if=%s", pgdata)
	dst := fmt.Sprintf("of=%s", target)
	bs := "bs=4096"
	out, err := exec.Command(dd, src, dst, bs).Output()
	log.Println(out)
	if err != nil {
		log.Fatalf("pgback dd %s %s: %s", src, dst, err)
	}
}

func Shutdown() {
	conn.Close()
}
