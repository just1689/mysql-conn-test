package main

import (
	"flag"
	"fmt"
	"github.com/just1689/mysql-conn-test/mct"
	"github.com/just1689/tracing"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	flag.Parse()
	traceId := tracing.NewId()

	//Setup Jaeger
	logrus.Println("setup tracing")
	tracing.StartTracing(tracing.Config{
		Url:             getEnvOr("tracingUrl", "http://jaeger-zipkin.cluster-infra.svc.cluster.local:8080/api/v2/spans"),
		CacheSize:       1024,
		FlushTimeout:    1,
		FlushSize:       1,
		SleepBetweenErr: 1,
	})

	logrus.Println("getting connection string")
	tracing.GlobalPublisher.Enqueue(tracing.NewSpan(traceId, mct.ServiceName, "Starting...", 1*time.Millisecond))
	connectionString := os.Getenv("connectionString")
	if connectionString == "" {
		time.Sleep(100 * time.Millisecond)
		tracing.GlobalPublisher.Enqueue(tracing.NewSpan(traceId, mct.ServiceName, "No connection string", 100*time.Millisecond))
	}

	logrus.Println("create config to connect to db")
	db, err := mct.TraceConnect(connectionString, traceId)
	if err != nil {
		time.Sleep(1 * time.Minute)
		panic(err)
	}

	logrus.Println("querying db")
	count, err := mct.QueryDateTraced(db, traceId)
	if err != nil {
		time.Sleep(1 * time.Minute)
		panic(err)
	}

	time.Sleep(100 * time.Millisecond)
	tracing.GlobalPublisher.Enqueue(tracing.NewSpan(traceId, mct.ServiceName, fmt.Sprint("Found rows", count), 100*time.Millisecond))

	time.Sleep(1 * time.Second)
}

func getEnvOr(name, def string) (r string) {
	r = os.Getenv(name)
	if r == "" {
		r = def
	}
	return
}
