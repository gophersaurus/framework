package bootstrap

import (
	"fmt"

	"github.com/gophersaurus/gf.v1/config"
	"github.com/gophersaurus/gf.v1/dba"
)

// DB bootstraps databases listed in configuration settings.
func DB() error {

	// find db config settings
	dbs := config.GetStringMap("databases")

	// each key value is the db name
	for name := range dbs {

		// find that particular db config map
		dmap := config.GetStringMapString("databases." + name)

		// create a new db instance based on db type
		switch dmap["type"] {
		case "mysql", "postgres", "sqlite":
			sql, err := dba.NewSQL(dmap["type"], dmap["user"], dmap["pass"], dmap["addr"])
			if err != nil {
				return err
			}
			dba.AddSQL(sql)
		case "mongo":
			mongo, err := dba.NewMongoDB(dmap["user"], dmap["pass"], dmap["addr"])
			if err != nil {
				return err
			}
			dba.AddMongoDB(mongo)
		default:
			return fmt.Errorf("unsupported %s database: %s", dmap["type"], name)
		}
	}

	return nil
}
