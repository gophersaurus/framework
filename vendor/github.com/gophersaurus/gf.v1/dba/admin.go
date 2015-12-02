// Package dba is an singleton package the manages databases for the
// gophersarus framework. DBA stands for Database Administrator.
package dba

import (
	"gopkg.in/gorp.v1"
	"gopkg.in/mgo.v2"
)

// admin is a database administrator object that manages different SQL and NoSQL
// databases. The admin object follows the singleton pattern in the db package
// so that admin is easy to access across a project.
var admin *DatabaseAdmin

func init() {
	admin = &DatabaseAdmin{}
}

// Database describes a methods to connect and close database connections.
type Database interface {
	Dial(name string) error
	Name() string
	Close()
}

// DatabaseAdmin represents a database administrator object who's duty is to
// maintain all database values, connections, and state.
type DatabaseAdmin struct {
	Mongo []Database
	SQL   []Database
}

// MGO takes a database name to search all mongodb databases db.DBA currently
// maintains and returns a pointer to a mgo.Database instance if found.
// mgo.Database represents a mongodb orm-ish driver for executing queries.
// If no database is found nil is returned.
func (d DatabaseAdmin) MGO(name string) *mgo.Database {
	for _, db := range d.Mongo {
		if db.Name() == name {
			if m, ok := db.(*MongoDB); ok {
				return m.mongodb
			}
			return nil
		}
	}
	return nil
}

// AddMongoDB adds a MongoDB database object to the DatabaseAdmin databases.
func (d *DatabaseAdmin) AddMongoDB(m *MongoDB) {
	d.Mongo = append(d.Mongo, m)
}

// GORP takes a database name to search all SQL databases DatabaseAdmin
// currently maintains and returns a pointer to a gorp.DbMap instance if found.
// gorp.DbMap represents an SQL orm-ish driver for multiple different SQL
// databases (sqlite, mysql, postgres). If no database is found nil is returned.
func (d DatabaseAdmin) GORP(name string) *gorp.DbMap {
	for _, db := range d.SQL {
		if db.Name() == name {
			if m, ok := db.(*SQL); ok {
				return m.sql
			}
			return nil
		}
	}
	return nil
}

// AddSQL adds a sql database object to the DatabaseAdmin databases.
func (d *DatabaseAdmin) AddSQL(s *SQL) {
	d.SQL = append(d.SQL, s)
}

// All returns all the databases this DatabaseAdmin maintains.
func (d *DatabaseAdmin) All() []Database {
	return append(d.Mongo, d.SQL...)
}

// MGO takes a database name to search all mongodb databases the db package
// currently maintains and returns a pointer to a mgo.Database instance if found.
// mgo.Database represents a mongodb orm-ish driver for executing queries.
// If no database is found nil is returned.
func MGO(name string) *mgo.Database {
	return admin.MGO(name)
}

// AddMongoDB adds a MongoDB database object to the db package.
func AddMongoDB(m *MongoDB) {
	admin.AddMongoDB(m)
}

// GORP takes a database name to search all SQL databases the db package
// currently maintains and returns a pointer to a gorp.DbMap instance if found.
// gorp.DbMap represents an SQL orm-ish driver for multiple different SQL
// databases (sqlite, mysql, postgres). If no database is found nil is returned.
func GORP(name string) *gorp.DbMap {
	return admin.GORP(name)
}

// AddSQL adds a sql database object to the singleton DatabaseAdmin's databases.
func AddSQL(s *SQL) {
	admin.AddSQL(s)
}

// All returns all the databases this singleton DatabaseAdmin maintains.
func All() []Database {
	return admin.All()
}
