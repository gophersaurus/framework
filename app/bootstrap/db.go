package bootstrap

import (
	"log"

	"github.com/gophersaurus/gf.v1/config"
	"github.com/gophersaurus/gf.v1/db"
)

// DB bootstraps databases listed in configuration settings.
func DB() error {

	dbs := config.GetStringMapString("databases")

	for _, d := range dbs {

		switch d.Type {
		case "mysql", "postgres", "sqlite":

			s, err := db.NewSQL(d.User, d.Pass, d.Address)
			if err != nil {
				return err
			}

			if err := s.Dial(d.Name, d.Type); err != nil {
				return err
			}

			db.AddSQL(s)

		case "mongo":

			m, err := db.NewMongoDB(d.User, d.Pass, d.Address)
			if err != nil {
				return err
			}

			if err := m.Dial(d.Name); err != nil {
				return err
			}

			db.AddMongoDB(m)

		default:
			log.Fatalf("unsupported %s database: %s", d.Type, d.Name)
		}
	}

	return nil
}
