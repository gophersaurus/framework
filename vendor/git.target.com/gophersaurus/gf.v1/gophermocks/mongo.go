package gophermocks

import (
	"net"
	"time"

	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mockstar"
	"git.target.com/gophersaurus/gophersaurus/vendor/git.target.com/gophersaurus/gf.v1/mongo"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2"
	"git.target.com/gophersaurus/gophersaurus/vendor/gopkg.in/mgo.v2/bson"
)

var (
	cast = &caster{}
)

type mockBulk struct {
	*mockstar.Mock
}

func NewMockBulk() *mockBulk {
	return &mockBulk{mockstar.NewMock()}
}

func (m *mockBulk) Insert(docs ...interface{}) {
	m.Mock.CalledVarArgs("Insert", docs)
}

func (m *mockBulk) Run() (*mgo.BulkResult, error) {
	args := m.Mock.Called("Run")
	return cast.asBulkResult(args.Get(0)), args.Err(1)
}

func (m *mockBulk) Unordered() {
	m.Mock.Called("Unordered")
}

type mockCollection struct {
	*mockstar.Mock
}

func NewMockCollection() *mockCollection {
	return &mockCollection{mockstar.NewMock()}
}

func (m *mockCollection) Database() mongo.Database {
	// TODO
	return nil
}

func (m *mockCollection) Name() string {
	// TODO
	return ""
}

func (m *mockCollection) FullName() string {
	// TODO
	return ""
}

func (m *mockCollection) Bulk() mongo.Bulk {
	args := m.Mock.Called("Bulk")
	return cast.asBulk(args.Get(0))
}

func (m *mockCollection) Count() (n int, err error) {
	args := m.Mock.Called("Count")
	return args.Int(0), args.Err(1)
}

func (m *mockCollection) Create(info *mgo.CollectionInfo) error {
	// TODO
	return nil
}

func (m *mockCollection) DropCollection() error {
	// TODO
	return nil
}

func (m *mockCollection) DropIndex(key ...string) error {
	// TODO
	return nil
}

func (m *mockCollection) EnsureIndex(index mgo.Index) error {
	// TODO
	return nil
}

func (m *mockCollection) EnsureIndexKey(key ...string) error {
	// TODO
	return nil
}

func (m *mockCollection) Find(query interface{}) mongo.Query {
	args := m.Mock.Called("Find", query)
	return cast.asQuery(args.Get(0))
}

func (m *mockCollection) FindId(id interface{}) mongo.Query {
	// TODO
	return nil
}

func (m *mockCollection) Indexes() (indexes []mgo.Index, err error) {
	// TODO
	return nil, nil
}

func (m *mockCollection) Insert(docs ...interface{}) error {
	// TODO
	return nil
}

func (m *mockCollection) Pipe(pipeline interface{}) mongo.Pipe {
	// TODO
	return nil
}

func (m *mockCollection) Remove(selector interface{}) error {
	// TODO
	return nil
}

func (m *mockCollection) RemoveAll(selector interface{}) (info *mgo.ChangeInfo, err error) {
	// TODO
	return nil, nil
}

func (m *mockCollection) RemoveId(id interface{}) error {
	args := m.Mock.Called("RemoveId", id)
	return args.Err(0)
}

func (m *mockCollection) Update(selector interface{}, update interface{}) error {
	// TODO
	return nil
}

func (m *mockCollection) UpdateAll(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	// TODO
	return nil, nil
}

func (m *mockCollection) UpdateId(id interface{}, update interface{}) error {
	// TODO
	return nil
}

func (m *mockCollection) Upsert(selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	// TODO
	return nil, nil
}

func (m *mockCollection) UpsertId(id interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	args := m.Mock.Called("UpsertId", id, update)
	return cast.asChangeInfo(args.Get(0)), args.Err(0)
}

func (m *mockCollection) With(s mongo.Session) mongo.Collection {
	// TODO
	return nil
}

type mockDatabase struct {
	*mockstar.Mock
}

func NewMockDatabase() *mockDatabase {
	return &mockDatabase{mockstar.NewMock()}
}

func (m *mockDatabase) Session() mongo.Session {
	// TODO
	return nil
}

func (m *mockDatabase) Name() string {
	// TODO
	return ""
}

func (m *mockDatabase) AddUser(username, password string, readOnly bool) error {
	// TODO
	return nil
}

func (m *mockDatabase) C(name string) mongo.Collection {
	args := m.Mock.Called("C", name)
	return cast.asCollection(args.Get(0))
}

func (m *mockDatabase) CollectionNames() (names []string, err error) {
	// TODO
	return nil, nil
}

func (m *mockDatabase) DropDatabase() error {
	// TODO
	return nil
}

func (m *mockDatabase) FindRef(ref *mgo.DBRef) mongo.Query {
	// TODO
	return nil
}

func (m *mockDatabase) GridFS(prefix string) mongo.GridFS {
	// TODO
	return nil
}

func (m *mockDatabase) Login(user, pass string) error {
	// TODO
	return nil
}

func (m *mockDatabase) Logout() {
	// TODO
}

func (m *mockDatabase) RemoveUser(user string) error {
	// TODO
	return nil
}

func (m *mockDatabase) Run(cmd interface{}, result interface{}) error {
	// TODO
	return nil
}

func (m *mockDatabase) UpsertUser(user *mgo.User) error {
	// TODO
	return nil
}

func (m *mockDatabase) With(s mongo.Session) mongo.Database {
	// TODO
	return nil
}

type mockGridFS struct {
	*mockstar.Mock
}

func NewMockGridFS() *mockGridFS {
	return &mockGridFS{mockstar.NewMock()}
}

func (m *mockGridFS) Create(name string) (file mongo.GridFile, err error) {
	// TODO
	return nil, nil
}

func (m *mockGridFS) Find(query interface{}) mongo.Query {
	// TODO
	return nil
}

func (m *mockGridFS) Open(name string) (file mongo.GridFile, err error) {
	// TODO
	return nil, nil
}

func (m *mockGridFS) OpenId(id interface{}) (file mongo.GridFile, err error) {
	// TODO
	return nil, nil
}

func (m *mockGridFS) OpenNext(iter mongo.Iter, file *mongo.GridFile) bool {
	// TODO
	return false
}

func (m *mockGridFS) Remove(name string) (err error) {
	// TODO
	return nil
}

func (m *mockGridFS) RemoveId(id interface{}) error {
	// TODO
	return nil
}

type mockGridFile struct {
	*mockstar.Mock
}

func NewMockGridFile() *mockGridFile {
	return &mockGridFile{mockstar.NewMock()}
}

func (m *mockGridFile) Abort() {
	// TODO
}

func (m *mockGridFile) Close() (err error) {
	// TODO
	return nil
}

func (m *mockGridFile) ContentType() string {
	// TODO
	return ""
}

func (m *mockGridFile) GetMeta(result interface{}) (err error) {
	// TODO
	return nil
}

func (m *mockGridFile) Id() interface{} {
	// TODO
	return nil
}

func (m *mockGridFile) MD5() (md5 string) {
	// TODO
	return ""
}

func (m *mockGridFile) Name() string {
	// TODO
	return ""
}

func (m *mockGridFile) Read(b []byte) (n int, err error) {
	// TODO
	return 0, nil
}

func (m *mockGridFile) Seek(offset int64, whence int) (pos int64, err error) {
	// TODO
	return 0, nil
}

func (m *mockGridFile) SetChunkSize(bytes int) {
	// TODO
}

func (m *mockGridFile) SetContentType(ctype string) {
	// TODO
}

func (m *mockGridFile) SetId(id interface{}) {
	// TODO
}

func (m *mockGridFile) SetMeta(metadata interface{}) {
	// TODO
}

func (m *mockGridFile) SetName(name string) {
	// TODO
}

func (m *mockGridFile) SetUploadDate(t time.Time) {
	// TODO
}

func (m *mockGridFile) Size() (bytes int64) {
	// TODO
	return 0
}

func (m *mockGridFile) UploadDate() time.Time {
	// TODO
	return time.Now()
}

func (m *mockGridFile) Write(data []byte) (n int, err error) {
	// TODO
	return 0, nil
}

type mockIter struct {
	*mockstar.Mock
}

func NewMockIter() *mockIter {
	return &mockIter{mockstar.NewMock()}
}

func (m *mockIter) All(result interface{}) error {
	// TODO
	return nil
}

func (m *mockIter) Close() error {
	// TODO
	return nil
}

func (m *mockIter) Err() error {
	// TODO
	return nil
}

func (m *mockIter) For(result interface{}, f func() error) (err error) {
	// TODO
	return nil
}

func (m *mockIter) Next(result interface{}) bool {
	// TODO
	return false
}

func (m *mockIter) Timeout() bool {
	// TODO
	return false
}

type mockPipe struct {
	*mockstar.Mock
}

func NewMockPipe() *mockPipe {
	return &mockPipe{mockstar.NewMock()}
}

func (m *mockPipe) All(result interface{}) error {
	// TODO
	return nil
}

func (m *mockPipe) AllowDiskUse() mongo.Pipe {
	// TODO
	return nil
}

func (m *mockPipe) Batch(n int) mongo.Pipe {
	// TODO
	return nil
}

func (m *mockPipe) Explain(result interface{}) error {
	// TODO
	return nil
}

func (m *mockPipe) Iter() mongo.Iter {
	// TODO
	return nil
}

func (m *mockPipe) One(result interface{}) error {
	// TODO
	return nil
}

type mockQuery struct {
	*mockstar.Mock
}

func NewMockQuery() *mockQuery {
	return &mockQuery{mockstar.NewMock()}
}

func (m *mockQuery) All(result interface{}) error {
	args := m.Mock.Called("All", result)
	return args.Err(0)
}

func (m *mockQuery) Apply(change mgo.Change, result interface{}) (info *mgo.ChangeInfo, err error) {
	// TODO
	return nil, nil
}

func (m *mockQuery) Batch(n int) mongo.Query {
	// TODO
	return nil
}

func (m *mockQuery) Count() (n int, err error) {
	// TODO
	return 0, nil
}

func (m *mockQuery) Distinct(key string, result interface{}) error {
	// TODO
	return nil
}

func (m *mockQuery) Explain(result interface{}) error {
	// TODO
	return nil
}

func (m *mockQuery) For(result interface{}, f func() error) error {
	// TODO
	return nil
}

func (m *mockQuery) Hint(indexKey ...string) mongo.Query {
	// TODO
	return nil
}

func (m *mockQuery) Iter() mongo.Iter {
	// TODO
	return nil
}

func (m *mockQuery) Limit(n int) mongo.Query {
	// TODO
	return nil
}

func (m *mockQuery) LogReplay() mongo.Query {
	// TODO
	return nil
}

func (m *mockQuery) MapReduce(job *mgo.MapReduce, result interface{}) (info *mgo.MapReduceInfo, err error) {
	// TODO
	return nil, nil
}

func (m *mockQuery) One(result interface{}) (err error) {
	args := m.Mock.Called("One", result)
	return args.Err(0)
}

func (m *mockQuery) Prefetch(p float64) mongo.Query {
	// TODO
	return nil
}

func (m *mockQuery) Select(selector interface{}) mongo.Query {
	args := m.Mock.Called("Select", selector)
	return cast.asQuery(args.Get(0))
}

func (m *mockQuery) SetMaxScan(n int) mongo.Query {
	// TODO
	return nil
}

func (m *mockQuery) Skip(n int) mongo.Query {
	// TODO
	return nil
}

func (m *mockQuery) Snapshot() mongo.Query {
	// TODO
	return nil
}

func (m *mockQuery) Sort(fields ...string) mongo.Query {
	// TODO
	return nil
}

func (m *mockQuery) Tail(timeout time.Duration) mongo.Iter {
	// TODO
	return nil
}

type mockServerAddr struct {
	*mockstar.Mock
}

func NewMockServerAddr() *mockServerAddr {
	return &mockServerAddr{mockstar.NewMock()}
}

func (m *mockServerAddr) String() string {
	// TODO
	return ""
}

func (m *mockServerAddr) TCPAddr() *net.TCPAddr {
	// TODO
	return nil
}

type mockSession struct {
	*mockstar.Mock
}

func NewMockSession() *mockSession {
	return &mockSession{mockstar.NewMock()}
}

func (m *mockSession) BuildInfo() (info mgo.BuildInfo, err error) {
	// TODO
	return mgo.BuildInfo{}, nil
}

func (m *mockSession) Clone() mongo.Session {
	// TODO
	return nil
}

func (m *mockSession) Close() {
	// TODO
}

func (m *mockSession) Copy() mongo.Session {
	// TODO
	return nil
}

func (m *mockSession) DB(name string) mongo.Database {
	// TODO
	return nil
}

func (m *mockSession) DatabaseNames() (names []string, err error) {
	// TODO
	return nil, nil
}

func (m *mockSession) EnsureSafe(safe *mgo.Safe) {
	// TODO
}

func (m *mockSession) FindRef(ref *mgo.DBRef) mongo.Query {
	// TODO
	return nil
}

func (m *mockSession) Fsync(async bool) error {
	// TODO
	return nil
}

func (m *mockSession) FsyncLock() error {
	// TODO
	return nil
}

func (m *mockSession) FsyncUnlock() error {
	// TODO
	return nil
}

func (m *mockSession) LiveServers() (addrs []string) {
	// TODO
	return nil
}

func (m *mockSession) Login(cred *mgo.Credential) error {
	// TODO
	return nil
}

func (m *mockSession) LogoutAll() {
	// TODO
}

func (m *mockSession) Mode() mongo.Mode {
	// TODO
	return mongo.Mode(0)
}

func (m *mockSession) New() mongo.Session {
	// TODO
	return nil
}

func (m *mockSession) Ping() error {
	// TODO
	return nil
}

func (m *mockSession) Refresh() {
	// TODO
}

func (m *mockSession) ResetIndexCache() {
	// TODO
}

func (m *mockSession) Run(cmd interface{}, result interface{}) error {
	// TODO
	return nil
}

func (m *mockSession) Safe() (safe *mgo.Safe) {
	// TODO
	return nil
}

func (m *mockSession) SelectServers(tags ...bson.D) {
	// TODO
}

func (m *mockSession) SetBatch(n int) {
	// TODO
}

func (m *mockSession) SetCursorTimeout(d time.Duration) {
	// TODO
}

func (m *mockSession) SetMode(consistency mongo.Mode, refresh bool) {
	// TODO
}

func (m *mockSession) SetPoolLimit(limit int) {
	// TODO
}

func (m *mockSession) SetPrefetch(p float64) {
	// TODO
}

func (m *mockSession) SetSafe(safe *mgo.Safe) {
	// TODO
}

func (m *mockSession) SetSocketTimeout(d time.Duration) {
	// TODO
}

func (m *mockSession) SetSyncTimeout(d time.Duration) {
	// TODO
}
