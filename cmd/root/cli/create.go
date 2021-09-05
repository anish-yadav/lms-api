package cli

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/anish-yadav/lms-api/internal/pkg/user"
	"github.com/anish-yadav/lms-api/internal/util"
	"github.com/peterbourgon/ff/v3/ffcli"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func CreateAdminUser() *ffcli.Command {
	var (
		flagset = flag.NewFlagSet("lms create-user", flag.ExitOnError)
		typ     = flagset.String("type", "", "the type of user you want to create")
	)

	return &ffcli.Command{
		Name:       "create-user",
		ShortUsage: "cosign create-user create a user for given type",
		ShortHelp:  "Create and add a user to database",
		LongHelp: `
		Create a user of type "admin" :
		lms create-user -type admin
		`,
		FlagSet: flagset,
		Exec: func(ctx context.Context, args []string) error {
			if len(*typ) == 0 {
				return errors.New("type is required")
			}

			if *typ != "admin" {
				return fmt.Errorf("type %s not supported", *typ)
			}

			return createUser(*typ)
		},
	}
}

func createUser(typ string) error {
	var email string
	var pwd string
	var name string

	reader := bufio.NewReader(os.Stdin)

	validatePwd := func(input string) error {
		if len(input) < 6 {
			return errors.New("password must have more than 6 characters")
		}
		return nil
	}
	confirmPwd := func(input string) error {
		if input != pwd {
			return errors.New("password does not match")
		}
		return nil
	}
	fmt.Print("Name: ")
	name, _ = reader.ReadString('\n')
	name = strings.Trim(name, "\n")
	log.Debugf("name is %s", name)
	fmt.Print("Email: ")
	email, _ = reader.ReadString('\n')
	email = strings.Trim(email, "\n")

	log.Debugf("email is %s", email)
	if err := util.ValidEmail(email); err != nil {
		log.Debugf("%s", err.Error())
		return errors.New("invalid email entered")
	}

	fmt.Print("Password: ")
	pwd, _ = reader.ReadString('\n')
	pwd = strings.Trim(pwd, "\n")
	if err := validatePwd(pwd); err != nil {
		return errors.New("password must be 6 character long")
	}

	fmt.Print("Confirm Password: ")
	cnfPwd, _ := reader.ReadString('\n')
	cnfPwd = strings.Trim(cnfPwd, "\n")
	if err := confirmPwd(cnfPwd); err != nil {
		return errors.New("password mismatch")
	}

	usr := user.NewUser(name, email, pwd, typ)
	id, err := usr.AddToDB()
	if err != nil {
		return err
	}
	log.Infof("Id of user is %s", id)
	return nil
}
