package mct

import (
	"database/sql"
	"fmt"
	"github.com/just1689/tracing"
	"time"
)

func QueryDateTraced(db *sql.DB, traceId string) (count int, err error) {
	rows, err := db.Query("SELECT NOW()")
	if err != nil {
		time.Sleep(100 * time.Millisecond)
		tracing.GlobalPublisher.Enqueue(tracing.NewSpan(traceId, ServiceName, fmt.Sprint("Could use db: ", err.Error()), 100*time.Millisecond))
		return
	}
	count = 0
	for rows.Next() {
		count++
	}
	time.Sleep(100 * time.Millisecond)
	tracing.GlobalPublisher.Enqueue(tracing.NewSpan(traceId, ServiceName, fmt.Sprint("Got date!"), 100*time.Millisecond))
	return
}
