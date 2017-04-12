package primes

import (
	"database/sql"
	"fmt"
	"os"
	"strings"
	// go-sqlite3 ...
	_ "github.com/mattn/go-sqlite3"
)

var dbfile = "./known.db"

// OnSQLiteKnowledge ...
type OnSQLiteKnowledge struct {
	db *sql.DB
}

// Init ...
func (s *OnSQLiteKnowledge) Init() error {
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		return err
	}
	s.db = db
	if _, err := s.db.Exec("CREATE TABLE IF NOT EXISTS knowns(n int64 primary key)"); err != nil {
		return err
	}
	return nil
}

// Clear ...
func (s *OnSQLiteKnowledge) Clear() error {
	if err := os.Remove(dbfile); err != nil {
		return err
	}
	return s.Init()
}

// Know ...
func (s *OnSQLiteKnowledge) Know(n int64) *Primes {
	return nil
}

// Learn ...
func (s *OnSQLiteKnowledge) Learn(p *Primes) Knowledge {
	query := "REPLACE INTO knowns(n) VALUES %s"
	pool := []string{}
	for _, n := range p.list {
		pool = append(pool, fmt.Sprintf("(%d)", n))
	}
	s.db.Exec(fmt.Sprintf(query, strings.Join(pool, ",")))
	return s
}

// Persist ...
func (s *OnSQLiteKnowledge) Persist(i int64) *Primes {
	rows, err := s.db.Query("SELECT n FROM knowns WHERE n <= ?", i)
	if err != nil {
		return Until(2)
	}
	defer rows.Close()
	p := &Primes{target: i, dictionary: map[int64]bool{}, list: []int64{}}
	for rows.Next() {
		var n int64
		if err := rows.Scan(&n); err != nil {
			continue
		}
		p.list = append(p.list, n)
		p.dictionary[n] = true
	}
	return p
}

// Until ...
func (s *OnSQLiteKnowledge) Until(n int64) *Primes {
	base := s.Persist(n)

	i := base.target
	if i == 1 {
		i = 2
	}
	p := extends(base, n)

	for ; i <= p.target; i++ {
		if p.knows(i) {
			continue // needless to evaluate.
		}
		if p.canDevideByKnownPrimeNumbers(i) {
			continue // it's not prime number.
		}
		// it's prime number,
		// and multiples of this number are no longer needless to be eveluated
		p.add(i)
	}

	s.Learn(p)

	return p
}
