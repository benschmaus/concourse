package commands

import (
	"errors"
	"fmt"
	"github.com/concourse/concourse/fly/commands/internal/displayhelpers"
	"github.com/concourse/concourse/fly/rc"
	"github.com/concourse/concourse/fly/ui"
	"github.com/fatih/color"
	"os"
	"time"
)

const inputDateLayout = "2006-01-02"

type UsersCommand struct {
	Since string `long:"since" description:""`
	Json bool `long:"json" description:"Print command result as JSON"`
}

func (command *UsersCommand) Execute([]string) error {
	target, err := rc.LoadTarget(Fly.Target, Fly.Verbose)
	if err != nil {
		return err
	}

	err = target.Validate()
	if err != nil {
		return err
	}

	dateSince := time.Now().AddDate(0, -2, 0)

	fmt.Println(dateSince)
	if command.Since != "" {
		dateSince, err = time.ParseInLocation(inputDateLayout, command.Since, time.Now().Location())
		if err != nil {
			return errors.New("since time should be in the format: " + inputTimeLayout)
		}
	}
	fmt.Println(dateSince)

	users, err := target.Client().ListActiveUsersSince(dateSince)
	fmt.Println(dateSince)
	if err != nil {
		return err
	}

	if command.Json {
		err = displayhelpers.JsonPrint(users)
		if err != nil {
			return err
		}
		return nil
	}

	headers := ui.TableRow{
		{Contents: "username", Color: color.New(color.Bold)},
		{Contents: "connector", Color: color.New(color.Bold)},
		{Contents: "last login", Color: color.New(color.Bold)},
	}

	table := ui.Table{Headers: headers}

	for _, user := range users {
		row := ui.TableRow{
			{Contents: user.Username},
			{Contents: user.Connector},
			{Contents: user.LastLogin.String()},
		}
		table.Data = append(table.Data, row)
	}
	return table.Render(os.Stdout, Fly.PrintTableHeaders)
}