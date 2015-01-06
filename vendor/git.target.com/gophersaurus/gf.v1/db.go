package gf

import "git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mongo"

var (
	Mgo  mongo.Database
	conn mongo.Session
)

func ConnectDB(conf *DbConfig) {
	defer FatalShock()
	dial := "mongodb://" + conf.Username + ":" + conf.Password + "@" + conf.Addr

	// connect to mongo database
	var err error
	conn, err = mongo.NewDialer().Dial(dial)
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
