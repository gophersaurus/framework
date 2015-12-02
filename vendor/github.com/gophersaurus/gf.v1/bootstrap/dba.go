package bootstrap

import (
	"fmt"

	"github.com/gophersaurus/gf.v1/config"
	"github.com/gophersaurus/gf.v1/dba"
)

// DBA bootstraps the dba object and dials the specified databases.
func DBA() error {

	// load database settings
	dbs := config.GetStringMap("databases")
	for name := range dbs {

		// find a particular db in the config map
		dmap := config.GetStringMapString("databases." + name)

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
