package mct

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/just1689/tracing"
	"github.com/sirupsen/logrus"
	"time"
)

func TraceConnect(c, traceId string) (db *sql.DB, err error) {
	db, err = sql.Open("mysql", c)
	if err != nil {
		logrus.Errorln("could not create connect config to db")
		time.Sleep(1 * time.Second)
		tracing.GlobalPublisher.Enqueue(tracing.NewSpan(traceId, ServiceName, err.Error(), 1*time.Second))
		return
	}
	time.Sleep(1 * time.Second)
	logrus.Println("created db config ok")
	tracing.GlobalPublisher.Enqueue(tracing.NewSpan(traceId, ServiceName, "db config ok", 1*time.Second))

	return
}

/*
	root:my@tcp(192.168.88.26)/my
*/
