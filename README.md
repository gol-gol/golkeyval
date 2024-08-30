
## golkeyval

> helper KeyVal integrations
>
> extracted from an older pandora box of such packages at [abhishekkr/gol](https://github.com/abhishekkr/gol)

Currently provides: `leveldb`, `badger`, `bitcask`, `sqlite3` & dumb `in-mem` via map

### Public Functions

* `golkeyval.GetDBEngine(dbtype)`: to init dbengine of above types providing type via values mentioned above

* `golkeyval.Configure(cfg map[string]string)`: to configure DB specific fields
> * badger: mandatory is `DBPath`; optional are `DetectConflicts:false`, `NumGoroutines`, `LogLevel:INFO`
> * bitcask: mandatory is `DBPath`
> * leveldb: mandatory is `DBPath`
> * sqlite3: mandatory is `DBPath`; optional is `TableName:golkeyval`

* `golkeyval.CreateDB()` to create required DB & get connection if required

* `golkeyval.CloseDB()` to cleanly close DB connection

* `golkeyval.CloseAndDeleteDB()` to close DB connection and purge its persistence

* `golkeyval.PushKeyVal(key string, val string) bool` to push a key-val

* `golkeyval.GetVal(key string) string` to fetch a val for key

* `golkeyval.DelKey(key string) bool` to remove a key-val

---
