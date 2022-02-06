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
