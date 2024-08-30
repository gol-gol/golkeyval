package golkeyval

import (
	"fmt"
	"os"

	pebbledb "github.com/cockroachdb/pebble"

	golerror "github.com/abhishekkr/gol/golerror"
)

/*
Pebble struct with required pebble details.
*/
type Pebble struct {
	DBPath string
	GolDB  *pebbledb.DB
}

/*
init registers pebble (GO implemented Pebble) to DBEngines.
*/
func init() {
	RegisterDBEngine("pebble", new(Pebble))
}

/*
Configure populates Pebble required configs.
*/
func (pbl *Pebble) Configure(cfg map[string]string) {
	pbl.DBPath = cfg["DBPath"]
}

/*
CreateDB creates a pebble db at provided DBPath.
*/
func (pbl *Pebble) CreateDB() {
	var errDB error
	opts := &pebbledb.Options{}
	opts = opts.EnsureDefaults()

	pbl.GolDB, errDB = pebbledb.Open(pbl.DBPath, opts)
	if errDB != nil {
		errMsg := fmt.Sprintf("DB %s Creation failed. %q", pbl.DBPath, errDB)
		golerror.Boohoo(errMsg, true)
	}
}

/*
CloseDB closes a db given handle.
*/
func (pbl *Pebble) CloseDB() {
	pbl.GolDB.Close()
}

/*
CloseAndDeleteDB closes and deletes a db given handle and DBPath.
Useful in use and throw implementations. And also tests.
*/
func (pbl *Pebble) CloseAndDeleteDB() {
	pbl.CloseDB()
	if os.RemoveAll(pbl.DBPath) != nil {
		panic("Fail: Temporary DB files are still present at: " + pbl.DBPath)
	}
}

/*
PushKeyVal pushes key-val in provided DB handle.
*/
func (pbl *Pebble) PushKeyVal(key string, val string) bool {
	if err := pbl.GolDB.Set([]byte(key), []byte(val), pebbledb.Sync); err != nil {
		golerror.Boohoo("Key "+key+" insertion failed. It's value was "+val, false)
		return false
	}
	return true
}

/*
GetVal return string-ified value of key from provided db handle.
*/
func (pbl *Pebble) GetVal(key string) string {
	data, closer, err := pbl.GolDB.Get([]byte(key))
	if err != nil {
		golerror.Boohoo("Key "+key+" query failed.", false)
		return ""
	}
	if errCloser := closer.Close(); errCloser != nil {
		golerror.Boohoo("Key "+key+" GET query's Pebble Closer failed.", false)
	}
	return string(data)
}

/*
DelKey deletes key from provided DB handle.
*/
func (pbl *Pebble) DelKey(key string) bool {
	err := pbl.GolDB.Delete([]byte(key), pebbledb.Sync)
	if err != nil {
		golerror.Boohoo("Key "+key+" query failed.", false)
		return false
	}
	return true
}
