package gophersauras

import (
	mgo "git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2"
)

var (
	Mgo  *mgo.Database
	conn *mgo.Session
)

func ConnectDB(conf *DbConfig) {
	defer FatalShock()
	dial := "mongodb://" + conf.Username + ":" + conf.Password + "@" + conf.Addr

	// connect to mongo database
	var err error
	conn, err = mgo.Dial(dial)
	Check(err)
	Mgo = conn.DB(conf.Name)
}

func CloseDB() {
	defer FatalShock()
	conn.Close()
}

func IsConnectionLost(err error) bool {
	return err != nil && err.Error() == "EOF"
}

func AttemptReconnect(retries, waitTime int) bool {
	// TODO --
	return false
}
