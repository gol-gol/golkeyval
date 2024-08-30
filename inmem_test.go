package golkeyval

import (
	"testing"

	"github.com/gol-gol/golassert"
)

func TestInMemInit(t *testing.T) {
	golassert.Type(t, DBEngines["in-mem"], new(InMem))
}

func TestInMemConfigureCreateAndCloseDB(t *testing.T) {
	kv := new(InMem)
	kv.Configure(map[string]string{})
	kv.CreateDB()
	kv.CloseDB()
}

func TestInMemPushGetDelDB(t *testing.T) {
	kv := new(InMem)
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
