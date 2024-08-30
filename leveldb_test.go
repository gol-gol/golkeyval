package golkeyval

import (
	"testing"

	"github.com/gol-gol/golassert"
)

func TestLevelDBInit(t *testing.T) {
	golassert.Type(t, DBEngines["leveldb"], new(LevelDB))
}

func TestLevelDBConfigureCreateAndCloseDB(t *testing.T) {
	kv := new(LevelDB)
	kv.Configure(map[string]string{"DBPath": "/tmp/test-db-leveldb"})
	kv.CreateDB()
	kv.CloseDB()
}

func TestLevelDBPushGetDelDB(t *testing.T) {
	kv := new(LevelDB)
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
