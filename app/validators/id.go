package validators

import (
	gf "../../vendor/git.target.com/gospot/framework"
	"../../vendor/gopkg.in/mgo.v2/bson"
)

// gets the 'id' url path parameter as a Mongo Id
func objectId(r *gf.Request) (bson.ObjectId, error) {
	return gf.ToBsonId(r.Vars["id"])
}
