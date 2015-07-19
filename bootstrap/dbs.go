package bootstrap

import (
	"log"

	"github.com/gophersaurus/framework/config"
)

// DBs bootstraps databases by reading a config object.
func DBs(c config.Config) error {

	// Iterate through the databases provided.
	for _, db := range c.Databases {

		switch db.Type {
		case "mysql", "postgres", "sqlite":

			s, err := db.NewSQL(db.User, db.Pass, db.Address)
			if err != nil {
				return err
			}

			if err := s.Dial(db.Name, db.Type); err != nil {
				return err
			}

			db.Admin.SQL[db.Name] = s

		case "mongo":

			s, err := db.NewMGO(db.User, db.Pass, db.Address)
			if err != nil {
				return err
			}

			if err := s.Dial(db.Name); err != nil {
				return err
			}

			db.Admin.MGO[db.Name] = s

		default:
			log.Fatalf("unsupported database: %s", db.Type)
		}
	}

	return dba
}
