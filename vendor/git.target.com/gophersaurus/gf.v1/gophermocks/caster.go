package gophermocks

import (
	"fmt"

	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mockstar"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2"
)

type caster struct {
}

func (c *caster) asBulkResult(obj interface{}) *mgo.BulkResult {
	if obj == nil {
		return nil
	}
	out, ok := obj.(mgo.BulkResult)
	if !ok {
		mockstar.Fatal("cannot cast value to mgo.BulkResult: " + fmt.Sprintf("%v", obj))
	}
	return &out
}

func (c *caster) asBulk(obj interface{}) *mockBulk {
	if obj == nil {
		return nil
	}
	out, ok := obj.(mockBulk)
	if !ok {
		mockstar.Fatal("cannot cast value to mockBulk: " + fmt.Sprintf("%v", obj))
	}
	return &out
}

func (c *caster) asQuery(obj interface{}) *mockQuery {
	if obj == nil {
		return nil
	}
	out, ok := obj.(mockQuery)
	if !ok {
		mockstar.Fatal("cannot cast value to mockQuery: " + fmt.Sprintf("%v", obj))
	}
	return &out
}

func (c *caster) asChangeInfo(obj interface{}) *mgo.ChangeInfo {
	if obj == nil {
		return nil
	}
	out, ok := obj.(mgo.ChangeInfo)
	if !ok {
		mockstar.Fatal("cannot cast value to mgo.ChangeInfo: " + fmt.Sprintf("%v", obj))
	}
	return &out
}

func (c *caster) asCollection(obj interface{}) *mockCollection {
	if obj == nil {
		return nil
	}
	out, ok := obj.(mockCollection)
	if !ok {
		mockstar.Fatal("cannot cast value to mockCollection: " + fmt.Sprintf("%v", obj))
	}
	return &out
}
