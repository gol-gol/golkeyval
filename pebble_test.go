package golkeyval

import (
	"testing"

	"github.com/gol-gol/golassert"
)

func TestPebbleInit(t *testing.T) {
	golassert.Type(t, DBEngines["pebble"], new(Pebble))
}

func TestPebbleConfigureCreateAndCloseDB(t *testing.T) {
	kv := new(Pebble)
	kv.Configure(map[string]string{"DBPath": "/tmp/test-db-leveldb"})
	kv.CreateDB()
	kv.CloseDB()
}

func TestPebblePushGetDelDB(t *testing.T) {
	kv := new(Pebble)
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
