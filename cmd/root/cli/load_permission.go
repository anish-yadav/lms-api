package cli

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"github.com/anish-yadav/lms-api/internal/pkg/permission"
	"github.com/peterbourgon/ff/v3/ffcli"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

func LoadPermissions() *ffcli.Command {
	var (
		flagset = flag.NewFlagSet("lms load-permission", flag.ExitOnError)
		file    = flagset.String("file", "", "the file contains the permission map")
	)

	return &ffcli.Command{
		Name:       "load-permission",
		ShortUsage: "lms load-permission -file permission.json populates the permissions db",
		FlagSet:    flagset,
		Exec: func(ctx context.Context, args []string) error {
			if len(*file) == 0 {
				return errors.New("file path is required")
			}

			return populatePermissionTable(*file)
		},
	}
}

func populatePermissionTable(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Debugf("populatePermission: %s", err.Error())
		return errors.New("failed to read file")
	}
	if err = permission.ClearDB(); err != nil {
		log.Errorf("failed to clear db")
		return err
	}
	var permissionMap map[string][]string
	err = json.Unmarshal(data, &permissionMap)

	for name, permissions := range permissionMap {
		currPermission := permission.NewPermission(name, permissions)
		id, err := currPermission.AddToDB()
		if err != nil {
			return err
		}
		log.Infof("document inserted with id: %s", id)
	}
	return nil
}
