package bootstrap

import (
	"fmt"
	"log"

	"github.com/gophersaurus/framework/config"
	"github.com/gophersaurus/gf.v1"
)

// Databases takes a config object and returns a gf.DBA
func Databases(c config.Config) *gf.DBA {

	// Create a new DBA to work with.
	dba := gf.NewDBA()

	// Iterate through the databases provided.
	for _, db := range c.Databases {

		// Switch to find the dbs we support.
		switch db.Type {
		case "mysql":

			// Create a new gorp
			g, err := gf.NewGorp(db.User, db.Pass, db.Address)

			// Check for error.
			if err != nil {
				log.Fatal(err)
			}

			// Connect
			fmt.Print("	Attempting to connect to MySQL " + db.Name + "... ")
			if err := g.Connect("mysql", db.Name); err != nil {
				log.Fatalln("Failed: " + err.Error())
			}

			// Let the user know we have connected.
			fmt.Println("Success!")

			// Assign the gorp to its name in the DBA.
			dba.SQL[db.Name] = g

		case "mongo", "mongodb":

			// Create a new MongoDB
			m, err := gf.NewMongoDB(db.User, db.Pass, db.Address)

			// Check for error.
			if err != nil {
				log.Fatal(err)
			}

			// Connect
			if err := m.Connect(db.Name); err != nil {
				log.Fatalln("Failed: " + err.Error())
			}

			// Assign the mongodb to its name in the DBA.
			dba.NoSQL[db.Name] = m

		case "postgres", "postgresql":

			// Create a new gorp
			g, err := gf.NewGorp(db.User, db.Pass, db.Address)

			// Check for error.
			if err != nil {
				log.Fatal(err)
			}

			// Connect
			if err := g.Connect("postgres", db.Name); err != nil {
				log.Fatalln("Failed: " + err.Error())
			}

			// Assign the gorp to its name in the DBA.
			dba.SQL[db.Name] = g

		default:
			log.Fatalln("Unsupported database: " + db.Type)
		}
	}

	return dba
}
