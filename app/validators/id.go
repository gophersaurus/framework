package validators

import (
	gf "git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/framework"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2/bson"
)

// gets the 'id' url path parameter as a Mongo Id
func ObjectId(r *gf.Request) (bson.ObjectId, error) {
	return gf.StringToBsonId(r.Vars["id"])
}
