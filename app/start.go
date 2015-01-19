package app

import (
	"log"
	"net/http"
	"strconv"

	"git.target.com/gophersaurus/gf.v1"
)

func Start(keys gf.KeyMap, db *gf.DbConfig, port int, indentJson bool) {
	if db != nil && db.IsValid() {
		// connect to database
		gf.ConnectDB(db)

		// Defer the command to close the MongoDB connection.
		defer gf.CloseDB()
	}

	// Create a new router.
	r := gf.NewRouter()

	// register dynamic routes.
	register(keys, r)

	// Start the http server on the correct port.
	portStr := strconv.Itoa(port)
	log.Println("listening on port " + portStr)
	http.ListenAndServe(":"+portStr, r)
}
