package cli

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/anish-yadav/lms-api/internal/pkg/user"
	"github.com/peterbourgon/ff/v3/ffcli"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

func ChangePassword() *ffcli.Command {
	var (
		flagset = flag.NewFlagSet("lms change-password", flag.ExitOnError)
		email   = flagset.String("email", "", "the email of the user you want to reset")
	)

	return &ffcli.Command{
		Name:       "change-password",
		ShortUsage: "lms change-password changes password of the given user",
		ShortHelp:  "Change password for the given user",
		LongHelp: `
		Change password for the given user :
		lms change-password -email john.doe@gmail.com
		`,
		FlagSet: flagset,
		Exec: func(ctx context.Context, args []string) error {
			if len(*email) == 0 {
				return errors.New("enter a valid email address")
			}
			return changePassword(*email)
		},
	}
}

func changePassword(email string) error {
	var pwd string

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

	usr := user.GetUserByEmail(email)
	if usr == nil {
		return errors.New("no user found with email: " + email)
	}
	err := usr.ResetPassword(pwd)
	if err != nil {
		return err
	}
	log.Infof("password reset successful")
	return nil
}
