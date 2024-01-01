package tools


type LogDb interface {
	LogRequestToDB(Timestamp string, Query string)
	GetLogs() LogDatabase
}

func GetLogDB() *LogDb {
	
	var db LogDb = getClient()

	return &db

}