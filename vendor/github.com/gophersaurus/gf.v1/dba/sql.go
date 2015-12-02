package dba

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	_ "github.com/lib/pq"              // postgresql driver

	// TODO:(jackspirou) this sqlite driver is written in C, so it takes a very
	// long time to build.
	//
	// Since this driver is written in C, it also stops cross compiles... not cool.
	//
	// Eventually it should be replaced.
	//
	// _ "github.com/mattn/go-sqlite3" // sqlite driver

	"gopkg.in/gorp.v1"
)

// SQL describes an sql database object.
type SQL struct {
	sql    *gorp.DbMap
	driver string
	name   string
	user   string
	pass   string
	addr   string
}

// NewSQL creates a new sql database object.
func NewSQL(driver, user, pass, addr string) (*SQL, error) {

	if len(driver) < 1 {
		return nil, errors.New("a driver must be specified for a sql database")
	}

	if !sqlDriver(driver) {
		return nil, fmt.Errorf("sql driver %s is not suppported", driver)
	}

	if len(user) == 0 {
		return nil, fmt.Errorf("missing sql username for %s", addr)
	}

	if len(pass) == 0 {
		return nil, fmt.Errorf("missing sql password for %s", addr)
	}

	if len(addr) == 0 {
		return nil, errors.New("missing sql address ")
	}

	return &SQL{driver: driver, user: user, pass: pass, addr: addr}, nil
}

// Name satisfies the Database interface.  The sql database object returns
// the name of the last database the Dial method attempted to connect with.
func (s *SQL) Name() string {
	return s.name
}

// Dial satisfies the Database interface. The sql database object attempts
// to dial and start a session with a local or remote sql instance.
func (s *SQL) Dial(name string) error {

	// recover from any internal panics that might occur, and log them
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("error connecting to %s database %s: %s", s.driver, name, r)
		}
	}()

	// connection string
	dial := "user=" + s.user + " dbname=" + name + " host=" + s.addr + "password=" + s.pass

	// dial and establish database session
	db, err := sql.Open(s.driver, dial)
	if err != nil {
		return err
	}

	// set database session and dialect default settings
	s.sql = &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	return nil
}

// Close satisfies the Database interface.  The sql database object closes the
// connection the Dial method created.
func (s *SQL) Close() {

	// recover from any internal panics that might occur, and log them
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("error closing sql database %s: %s", s.name, r)
		}
	}()

	// close the connection.
	s.sql.Db.Close()
}

func sqlDriver(driver string) bool {
	return driver == "mysql" || driver == "postgres" || driver == "sqlite"
}
