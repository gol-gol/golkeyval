package golkeyval

import "testing"

/*
Defining a Fake DBEngine to get used across tests that require so.
*/
type FakeDBEngine struct{}

func (f *FakeDBEngine) Configure(cfg map[string]string) {
	return
}
func (f *FakeDBEngine) CreateDB() {
	return
}
func (f *FakeDBEngine) CloseDB() {
	return
}
func (f *FakeDBEngine) CloseAndDeleteDB() {
	return
}
func (f *FakeDBEngine) PushKeyVal(key string, val string) bool {
	return true
}
func (f *FakeDBEngine) GetVal(key string) string {
	if key == "fake-key" {
		return "fake-val"
	}
	return ""
}
func (f *FakeDBEngine) DelKey(key string) bool {
	return true
}

func TestRegisterDBEngine(t *testing.T) {
	RegisterDBEngine("fake-db", new(FakeDBEngine))
	if DBEngines["fake-db"].GetVal("fake-key") != "fake-val" {
		t.Fatal("RegisterDBEngine failed to call GetVal at FakeDBEngine.")
	}
}

func TestGetDBEngine(t *testing.T) {
	DBEngines["fake-db"] = new(FakeDBEngine)
	if GetDBEngine("fake-db").GetVal("fake-key") != "fake-val" {
		t.Fatal("GetDBEngine failed to call GetVal at FakeDBEngine.")
	}
}
