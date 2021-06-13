package db


type Db interface {
	InsertRecord(table string, id string, rowData map[string]interface{}) bool
	GetRecordById(table string, id string) interface{}
	Init()
}
