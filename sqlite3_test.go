package golkeyval

import (
	"os"
	"testing"

	"github.com/gol-gol/golassert"
)

func TestSqlite3DBInit(t *testing.T) {
	golassert.Type(t, DBEngines["sqlite3"], new(Sqlite3DB))
}

func TestSqlite3DBConfigureCreateAndCloseDB(t *testing.T) {
	kv := new(Sqlite3DB)
	kv.Configure(map[string]string{
		"DBPath":    "/tmp/test-db-sqlite3",
		"TableName": "testkv",
	})
	kv.CreateDB()
	kv.CloseDB()
	if os.RemoveAll(kv.DBPath) != nil {
		panic("Fail: Temporary DB files are still present at: " + kv.DBPath)
	}
}

func TestSqlite3DBPushGetDelDB(t *testing.T) {
	kv := new(Sqlite3DB)
	kv.Configure(map[string]string{"DBPath": "/tmp/test-db-leveldb"})
	kv.CreateDB()
	golassert.Equal(t, "", kv.GetVal("sample-key"))
	golassert.Equal(t, true, kv.PushKeyVal("sample-key", "sample-value"))
	golassert.Equal(t, "sample-value", kv.GetVal("sample-key"))
	golassert.Equal(t, true, kv.PushKeyVal("sample-key", "next-value"))
	golassert.Equal(t, "next-value", kv.GetVal("sample-key"))
	golassert.Equal(t, true, kv.DelKey("sample-key"))
	golassert.Equal(t, "", kv.GetVal("sample-key"))
	kv.CloseAndDeleteDB()
}
