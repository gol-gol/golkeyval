package golkeyval

type InMem struct {
	KeyVal map[string]string
}

/*
init registers InMem (a dumb map based key-val) to DBEngines.
*/
func init() {
	RegisterDBEngine("in-mem", new(InMem))
}

/*
Configure is for sake of interface only, placeholder.
*/
func (inMem *InMem) Configure(cfg map[string]string) {
	return
}

/*
CreateDB just initializes Map.
*/
func (inMem *InMem) CreateDB() {
	inMem.KeyVal = make(map[string]string)
}

/*
CloseDB deletes all keys from Map.
*/
func (inMem *InMem) CloseDB() {
	for k := range inMem.KeyVal {
		delete(inMem.KeyVal, k)
	}
}

/*
CloseAndDeleteDB just a proxy for CloseDB here.
*/
func (inMem *InMem) CloseAndDeleteDB() {
	inMem.CloseDB()
}

/*
PushKeyVal pushes key-val in provided inMem map.
*/
func (inMem *InMem) PushKeyVal(key string, val string) bool {
	inMem.KeyVal[key] = val
	return true
}

/*
GetVal return val for provided key in InMem map.
*/
func (inMem *InMem) GetVal(key string) string {
	return inMem.KeyVal[key]
}

/*
DelKey deletes key from map.
*/
func (inMem *InMem) DelKey(key string) bool {
	delete(inMem.KeyVal, key)
	return true
}
