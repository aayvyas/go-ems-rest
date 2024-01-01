package tools

import (
	"time"

	"github.com/sirupsen/logrus"
)
type LogEntry struct{
	Timestamp string
	Query string
}

type LogDatabase struct {
	Count int
	Entries []LogEntry
}
var logEntries []LogEntry = []LogEntry{
	{
		Timestamp: time.Now().String(),
		Query: "Some query",
	},
	{
		Timestamp: time.August.String(),
		Query: "Some other query",
	},
}

var logDB = LogDatabase{
	Entries: logEntries,
	Count: len(logEntries),
}

func (e *LogDatabase) LogRequestToDB(Timestamp string, Query string) {
	var logEntry = LogEntry{
		Timestamp:  Timestamp,
		Query: Query,
	}
	logDB.Entries = append(logDB.Entries, logEntry)
	logrus.Info(logDB)
}

func getClient() *LogDatabase{
	return &logDB
}

func (e *LogDatabase) GetLogs() LogDatabase {
	return logDB
}
