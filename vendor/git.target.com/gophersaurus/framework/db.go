package gophersauras

import mgo "git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2"

var (
	conn *mgo.Session
	Mgo  *mgo.Database
)

func ConnectDB(db *DbConfig) {
	dial := "mongodb://" + db.Username + ":" + db.Password + "@" + db.Addr

	// connect to mongo database
	var err error
	conn, err = mgo.Dial(dial)
	Check(err)
	Mgo = conn.DB(db.Name)
}

func CloseDB() {
	conn.Close()
}
