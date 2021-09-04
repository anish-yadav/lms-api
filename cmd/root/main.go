package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/anish-yadav/lms-api/cmd/root/cli"
	"github.com/anish-yadav/lms-api/internal/pkg/db"
	"github.com/peterbourgon/ff/v3/ffcli"
	"os"
)

var (
	dbURI  = flag.String("dbAddr", "mongodb://localhost:27017", "url of mongodb database")
	dbName = flag.String("db", "lms", "database name")
)

func main() {
	flag.Parse()
	db.Init(*dbURI, *dbName)

	var (
		rootFlagSet = flag.NewFlagSet("lms", flag.ExitOnError)
	)

	root := &ffcli.Command{
		ShortUsage:  "lms [flags] <subcommand>",
		FlagSet:     rootFlagSet,
		Subcommands: []*ffcli.Command{
			cli.CreateAdminUser(),
		},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		fmt.Printf("%s",err)
		return
	}
}
