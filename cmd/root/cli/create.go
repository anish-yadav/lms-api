package cli

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/anish-yadav/lms-api/internal/pkg/user"
	"github.com/anish-yadav/lms-api/internal/util"
	"github.com/manifoldco/promptui"
	"github.com/peterbourgon/ff/v3/ffcli"
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
	emailPrompt := promptui.Prompt{
		Label:    "Email",
		Validate: util.ValidEmail,
	}
	email, err := emailPrompt.Run()
	if err != nil {
		return fmt.Errorf("invalid email address")
	}

	pwdPrompt := promptui.Prompt{
		Label:    "Password",
		Validate: validatePwd,
		Mask:     '*',
	}
	pwd, err = pwdPrompt.Run()
	if err != nil {
		return fmt.Errorf("password must be 6 digit long")
	}

	cnfPwdPrompt := promptui.Prompt{
		Label:    "Confirm Password",
		Validate: confirmPwd,
		Mask:     '*',
	}
	pwd, err = cnfPwdPrompt.Run()
	if err != nil {
		return fmt.Errorf("password must be 6 digit long")
	}

	usr := user.NewUser(email, pwd, typ)
	return usr.AddToDB()
}
