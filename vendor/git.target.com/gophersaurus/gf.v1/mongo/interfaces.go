package mongo

import (
	"net"
	"time"

	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2/bson"
)

type Bulk interface {
	Insert(docs ...interface{})
	Run() (*mgo.BulkResult, error)
	Unordered()
}
type Collection interface {
	// properties
	Database() Database
	Name() string
	FullName() string

	// methods
	Bulk() Bulk
	Count() (n int, err error)
	Create(info *mgo.CollectionInfo) error
	DropCollection() error
	DropIndex(key ...string) error
	EnsureIndex(index mgo.Index) error
	EnsureIndexKey(key ...string) error
	Find(query interface{}) Query
	FindId(id interface{}) Query
	Indexes() (indexes []mgo.Index, err error)
	Insert(docs ...interface{}) error
	Pipe(pipeline interface{}) Pipe
	Remove(selector interface{}) error
	RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error)
	RemoveId(id interface{}) error
	//Repair() Iter
	Update(selector interface{}, update interface{}) error
	UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	UpdateId(id interface{}, update interface{}) error
	Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	UpsertId(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error)
	With(s Session) Collection
}
type Database interface {
	// Properties
	Session() Session
	Name() string

	// Methods
	AddUser(username, password string, readOnly bool) error
	C(name string) Collection
	CollectionNames() (names []string, err error)
	DropDatabase() error
	FindRef(ref *mgo.DBRef) Query
	GridFS(prefix string) GridFS
	Login(user, pass string) error
	Logout()
	RemoveUser(user string) error
	Run(cmd interface{}, result interface{}) error
	UpsertUser(user *mgo.User) error
	With(s Session) Database
}
type GridFS interface {
	//Properties
	Files() Collection
	Chunks() Collection

	//Methods
	Create(name string) (file GridFile, err error)
	Find(query interface{}) Query
	Open(name string) (file GridFile, err error)
	OpenId(id interface{}) (file GridFile, err error)
	OpenNext(iter Iter, file *GridFile) bool
	Remove(name string) (err error)
	RemoveId(id interface{}) error
}
type GridFile interface {
	Abort()
	Close() (err error)
	ContentType() string
	GetMeta(result interface{}) (err error)
	Id() interface{}
	MD5() (md5 string)
	Name() string
	Read(b []byte) (n int, err error)
	Seek(offset int64, whence int) (pos int64, err error)
	SetChunkSize(bytes int)
	SetContentType(ctype string)
	SetId(id interface{})
	SetMeta(metadata interface{})
	SetName(name string)
	SetUploadDate(t time.Time)
	Size() (bytes int64)
	UploadDate() time.Time
	Write(data []byte) (n int, err error)
}
type Iter interface {
	All(result interface{}) error
	Close() error
	Err() error
	For(result interface{}, f func() error) (err error)
	Next(result interface{}) bool
	Timeout() bool
}

type Pipe interface {
	All(result interface{}) error
	AllowDiskUse() Pipe
	Batch(n int) Pipe
	Explain(result interface{}) error
	Iter() Iter
	One(result interface{}) error
}

type Query interface {
	All(result interface{}) error
	Apply(change mgo.Change, result interface{}) (info *mgo.ChangeInfo, err error)
	Batch(n int) Query
	Count() (n int, err error)
	Distinct(key string, result interface{}) error
	Explain(result interface{}) error
	For(result interface{}, f func() error) error
	Hint(indexKey ...string) Query
	Iter() Iter
	Limit(n int) Query
	LogReplay() Query
	MapReduce(job *mgo.MapReduce, result interface{}) (info *mgo.MapReduceInfo, err error)
	One(result interface{}) (err error)
	Prefetch(p float64) Query
	Select(selector interface{}) Query
	SetMaxScan(n int) Query
	Skip(n int) Query
	Snapshot() Query
	Sort(fields ...string) Query
	Tail(timeout time.Duration) Iter
}

type ServerAddr interface {
	String() string
	TCPAddr() *net.TCPAddr
}

type Session interface {
	BuildInfo() (info mgo.BuildInfo, err error)
	Clone() Session
	Close()
	Copy() Session
	DB(name string) Database
	DatabaseNames() (names []string, err error)
	EnsureSafe(safe *mgo.Safe)
	FindRef(ref *mgo.DBRef) Query
	Fsync(async bool) error
	FsyncLock() error
	FsyncUnlock() error
	LiveServers() (addrs []string)
	Login(cred *mgo.Credential) error
	LogoutAll()
	Mode() Mode
	New() Session
	Ping() error
	Refresh()
	ResetIndexCache()
	Run(cmd interface{}, result interface{}) error
	Safe() (safe *mgo.Safe)
	SelectServers(tags ...bson.D)
	SetBatch(n int)
	SetCursorTimeout(d time.Duration)
	SetMode(consistency Mode, refresh bool)
	SetPoolLimit(limit int)
	SetPrefetch(p float64)
	SetSafe(safe *mgo.Safe)
	SetSocketTimeout(d time.Duration)
	SetSyncTimeout(d time.Duration)
}

type Mode int

const (
	Eventual  Mode = 0
	Monotonic Mode = 1
	Strong    Mode = 2
)

type Dialer interface {
	Dial(url string) (Session, error)
	DialWithInfo(info *mgo.DialInfo) (Session, error)
	DialWithTimeout(url string, timeout time.Duration) (Session, error)
}
