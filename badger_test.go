package golkeyval

import (
	"os"
	"testing"

	"github.com/gol-gol/golassert"
)

func TestBadgerInit(t *testing.T) {
	golassert.Type(t, DBEngines["badger"], new(Badger))
}

func TestBadgerConfigureCreateAndCloseDB(t *testing.T) {
	kv := new(Badger)
	kv.Configure(map[string]string{"DBPath": "/tmp/test-db-badger"})
	kv.CreateDB()
	kv.CloseDB()
	if os.RemoveAll(kv.DBPath) != nil {
		panic("Fail: Temporary DB files are still present at: " + kv.DBPath)
	}
}

func TestBadgerPushGetDelDB(t *testing.T) {
	kv := new(Badger)
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
