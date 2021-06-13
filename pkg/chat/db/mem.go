package db

type MemDataAccess struct {
	db map[string]map[string]interface{}
}

func (memDataAccess MemDataAccess) Init() error {
	return nil
}

func (memDataAccess MemDataAccess) InsertRecord(table string, id string, rowData map[string]interface{}) bool {
	if _, ok := memDataAccess.db[table]; !ok {
		memDataAccess.db[table] = make(map[string]interface{})
	}
	memDataAccess.db[table][id] = rowData
	return true
}

func (memDataAccess MemDataAccess) GetRecordById(table string, id string) interface{} {
	if _, ok := memDataAccess.db[table]; !ok {
		return nil
	}
	if _, ok := memDataAccess.db[table][id]; !ok {
		return nil
	}
	return memDataAccess.db[table][id]
}
