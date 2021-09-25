package main

import (
	"flag"
	"fmt"
	"github.com/anish-yadav/lms-api/internal/constants"
	"github.com/anish-yadav/lms-api/internal/pkg/db"
	"github.com/anish-yadav/lms-api/internal/pkg/webservice"
	"github.com/anish-yadav/lms-api/internal/util"
	"github.com/google/uuid"
	"os"
)

var (
	dbURI  = flag.String("dbAddr", "mongodb://localhost:27017", "url of mongodb database")
	dbName = flag.String("db", "lms", "database name")
	port   = flag.String("port", "8080", "port of the server")
	log    = flag.String("log", "debug", "log level")
)

func init() {
	jwtSecret := os.Getenv(constants.JwtSecret)
	if len(jwtSecret) == 0 {
		jwtSecret = uuid.New().String()
		fmt.Println(jwtSecret)
		os.Setenv(constants.JwtSecret, jwtSecret)
	}
}

func main() {

	flag.Parse()
	db.Init(*dbURI, *dbName)
	util.InitLogger(*log)

	webservice.StartServer(*port)
}
