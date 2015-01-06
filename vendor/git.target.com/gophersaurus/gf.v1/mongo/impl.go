package mongo

import (
	"time"

	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2"
)

func NewDialer() Dialer {
	return &dialer{}
}

type dialer struct {
}

func (d *dialer) Dial(url string) (Session, error) {
	sess, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	return &session{sess}, nil
}
func (d *dialer) DialWithInfo(info *mgo.DialInfo) (Session, error) {
	sess, err := mgo.DialWithInfo(info)
	if err != nil {
		return nil, err
	}
	return &session{sess}, nil
}
func (d *dialer) DialWithTimeout(url string, timeout time.Duration) (Session, error) {
	sess, err := mgo.DialWithTimeout(url, timeout)
	if err != nil {
		return nil, err
	}
	return &session{sess}, nil
}

type bulk struct {
	*mgo.Bulk
}

type collection struct {
	*mgo.Collection
}

func (c *collection) Database() Database {
	return &database{c.Collection.Database}
}

func (c *collection) Name() string {
	return c.Collection.Name
}

func (c *collection) FullName() string {
	return c.Collection.FullName
}

func (c *collection) Bulk() Bulk {
	return &bulk{c.Collection.Bulk()}
}

func (c *collection) Find(q interface{}) Query {
	return &query{c.Collection.Find(q)}
}

func (c *collection) FindId(id interface{}) Query {
	return &query{c.Collection.FindId(id)}
}

func (c *collection) Pipe(pipeline interface{}) Pipe {
	return &pipe{c.Collection.Pipe(pipeline)}
}

/*
func (c *collection) Repair() Iter {
	return &iter{c.Collection.Repair()}
}
*/

func (c *collection) With(s Session) Collection {
	// TODO -- session type checking
	return &collection{ /*collection instance*/ }
}

type database struct {
	*mgo.Database
}

func (d *database) Session() Session {
	return &session{d.Database.Session}
}

func (d *database) Name() string {
	return d.Database.Name
}

func (d *database) C(name string) Collection {
	return &collection{d.Database.C(name)}
}

func (d *database) FindRef(ref *mgo.DBRef) Query {
	return &query{d.Database.FindRef(ref)}
}

func (d *database) GridFS(prefix string) GridFS {
	return &gridfs{d.Database.GridFS(prefix)}
}

func (d *database) With(s Session) Database {
	// TODO -- session type checking
	return &database{ /*database instance*/ }
}

type gridfs struct {
	*mgo.GridFS
}

func (g *gridfs) Files() Collection {
	return &collection{g.GridFS.Files}
}
func (g *gridfs) Chunks() Collection {
	return &collection{g.GridFS.Chunks}
}

func (g *gridfs) Create(name string) (file GridFile, err error) {
	gf, err := g.GridFS.Create(name)
	if err != nil {
		return nil, err
	}
	return &gridfile{gf}, nil
}

func (g *gridfs) Find(q interface{}) Query {
	return &query{g.GridFS.Find(q)}
}

func (g *gridfs) Open(name string) (file GridFile, err error) {
	gf, err := g.GridFS.Open(name)
	if err != nil {
		return nil, err
	}
	return &gridfile{gf}, nil
}

func (g *gridfs) OpenId(id interface{}) (file GridFile, err error) {
	gf, err := g.GridFS.OpenId(id)
	if err != nil {
		return nil, err
	}
	return &gridfile{gf}, nil
}

func (g *gridfs) OpenNext(iter Iter, file *GridFile) bool {
	// TODO -- type checking
	return false /*database instance*/
}

type gridfile struct {
	*mgo.GridFile
}

type iter struct {
	*mgo.Iter
}

type pipe struct {
	*mgo.Pipe
}

func (p *pipe) AllowDiskUse() Pipe {
	return &pipe{p.Pipe.AllowDiskUse()}
}

func (p *pipe) Batch(n int) Pipe {
	return &pipe{p.Pipe.Batch(n)}
}

func (p *pipe) Iter() Iter {
	return &iter{p.Pipe.Iter()}
}

type query struct {
	*mgo.Query
}

func (q *query) Batch(n int) Query {
	return &query{q.Query.Batch(n)}
}

func (q *query) Hint(indexKey ...string) Query {
	return &query{q.Query.Hint(indexKey...)}
}

func (q *query) Iter() Iter {
	return &iter{q.Query.Iter()}
}

func (q *query) Limit(n int) Query {
	return &query{q.Query.Limit(n)}
}

func (q *query) LogReplay() Query {
	return &query{q.Query.LogReplay()}
}

func (q *query) Prefetch(p float64) Query {
	return &query{q.Query.Prefetch(p)}
}

func (q *query) Select(selector interface{}) Query {
	return &query{q.Query.Select(selector)}
}

func (q *query) SetMaxScan(n int) Query {
	return &query{q.Query.SetMaxScan(n)}
}

func (q *query) Skip(n int) Query {
	return &query{q.Query.Skip(n)}
}

func (q *query) Snapshot() Query {
	return &query{q.Query.Snapshot()}
}

func (q *query) Sort(fields ...string) Query {
	return &query{q.Query.Sort(fields...)}
}

func (q *query) Tail(timeout time.Duration) Iter {
	return &iter{q.Query.Tail(timeout)}
}

type serverAddr struct {
	*mgo.ServerAddr
}

type session struct {
	*mgo.Session
}

func (s *session) Clone() Session {
	return &session{s.Session.Clone()}
}

func (s *session) Copy() Session {
	return &session{s.Session.Copy()}
}

func (s *session) DB(name string) Database {
	return &database{s.Session.DB(name)}
}

func (s *session) FindRef(ref *mgo.DBRef) Query {
	return &query{s.Session.FindRef(ref)}
}

func (s *session) Mode() Mode {
	out := s.Session.Mode()
	return Mode(int(out))
}

func (s *session) New() Session {
	return &session{s.Session.New()}
}

func (s *session) SetMode(consistency Mode, refresh bool) {
	inMode := mgo.Eventual
	switch consistency {
	case Eventual:
		inMode = mgo.Eventual
	case Monotonic:
		inMode = mgo.Monotonic
	case Strong:
		inMode = mgo.Strong
	}
	s.Session.SetMode(inMode, refresh)
}
