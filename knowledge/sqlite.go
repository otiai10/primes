package knowledge

import (
	"database/sql"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/otiai10/primes"
)

const (
	driver  = "sqlite3"
	rootdir = ".primes"
	dbname  = "known.db"
)

var q = map[string]string{
	"create_table_if_not_exists": `
    CREATE TABLE IF NOT EXISTS primes (
      number INTEGER PRIMARY KEY
    )
  `,
	"find_until": `SELECT number FROM primes WHERE number < ?`,
	// TODO: insert bulk
	"insert": `INSERT OR IGNORE INTO primes (number) VALUES %s`,
}

// SQLiteKnowledge ...
type SQLiteKnowledge struct {
	dir  string
	file string
	db   *sql.DB
}

// NewSQLiteKnowledge ...
func NewSQLiteKnowledge(dirname string) (*SQLiteKnowledge, error) {

	// Get user $HOME
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	// Define working directory path
	var dir string
	if dirname != "" {
		dir = filepath.Join(u.HomeDir, dirname)
	} else {
		dir = filepath.Join(u.HomeDir, rootdir)
	}

	// Create working directory
	_, err = os.Stat(dir)
	if os.IsNotExist(err) {
		if er := os.MkdirAll(dir, os.ModePerm); er != nil {
			return nil, er
		}
	} else if err != nil {
		return nil, err
	}

	// Create database file
	f := filepath.Join(dir, dbname)
	if _, er := os.Create(f); er != nil {
		return nil, err
	}

	// Create table if not exists
	db, err := sql.Open(driver, f)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	if _, err := db.Exec(q["create_table_if_not_exists"]); err != nil {
		return nil, err
	}

	return &SQLiteKnowledge{
		dir:  dir,
		file: f,
	}, nil
}

// initialize creats ~/.primes directory for database
func initialize() error {

	return nil
}

// Open opens db connection
func (k *SQLiteKnowledge) Open() error {
	db, err := sql.Open(driver, k.file)
	if err != nil {
		return err
	}
	k.db = db
	return nil
}

// Close closes db connection, it should be defered.
func (k *SQLiteKnowledge) Close() {
	k.db.Close()
}

// Know ...
func (k *SQLiteKnowledge) Know(target int64) *primes.Primes {
	rows, err := k.db.Query(q["find_until"], target)
	if err != nil {
		return nil
	}
	defer rows.Close()
	list := []int64{}
	for rows.Next() {
		var number int64
		err = rows.Scan(&number)
		if err != nil {
			return nil
		}
		// TODO: refactor by [len]int64{}
		list = append(list, number)
	}
	if len(list) == 0 {
		return nil
	}
	return primes.NewPrimes(target, list)
}

// Learn ...
func (k *SQLiteKnowledge) Learn(p *primes.Primes) {
	pool := []string{}
	for _, n := range p.List() {
		pool = append(pool, fmt.Sprintf("(%d)", n))
	}
	if _, err := k.db.Exec(fmt.Sprintf(q["insert"], strings.Join(pool, ","))); err != nil {
		panic(err)
	}
}
