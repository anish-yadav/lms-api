package main

import (
	"flag"
	"github.com/anish-yadav/lms-api/internal/pkg/db"
	"github.com/anish-yadav/lms-api/internal/pkg/webservice"
	"github.com/anish-yadav/lms-api/internal/util"
)

var (
	dbURI  = flag.String("dbAddr", "mongodb://localhost:27017", "url of mongodb database")
	dbName = flag.String("db", "lms", "database name")
	port   = flag.String("port", "8080", "port of the server")
	log    = flag.String("log", "debug", "log level")
)

func main() {

	flag.Parse()
	db.Init(*dbURI, *dbName)
	util.InitLogger(*log)


	webservice.StartServer(*port)
}
