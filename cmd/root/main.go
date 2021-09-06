package main

import (
	"context"
	"flag"
	"github.com/anish-yadav/lms-api/cmd/root/cli"
	"github.com/anish-yadav/lms-api/internal/pkg/db"
	"github.com/anish-yadav/lms-api/internal/util"
	"github.com/peterbourgon/ff/v3/ffcli"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	dbURI    = flag.String("dbAddr", "mongodb://localhost:27017", "url of mongodb database")
	dbName   = flag.String("db", "lms", "database name")
	logLevel = flag.String("log", "debug", "log level")
)

func main() {
	flag.Parse()
	db.Init(*dbURI, *dbName)
	db.CreateIndexes("users")
	util.InitLogger(*logLevel)

	var (
		rootFlagSet = flag.NewFlagSet("lms", flag.ExitOnError)
	)

	root := &ffcli.Command{
		ShortUsage: "lms [flags] <subcommand>",
		FlagSet:    rootFlagSet,
		Subcommands: []*ffcli.Command{
			cli.CreateAdminUser(),
			cli.LoadPermissions(),
		},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		log.Error(err)
		return
	}
}
