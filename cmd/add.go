package cmd

import (
	"fmt"
	"strings"

	"github.com/leoviggiano/daily/pkg/daily"
	"github.com/leoviggiano/daily/pkg/errors"
)

func Add(currentDaily *daily.Daily) Command {
	return &add{currentDaily: currentDaily}
}

type add struct {
	currentDaily *daily.Daily
}

func (a *add) Name() []string {
	return []string{"add", "a"}
}

func (a *add) Help() string {
	return fmt.Sprintf(`
	Adiciona um item a daily atual: %s

	Usage:
		add <item>

	Examples:
		add "Migrate genova template"
		add "Add new response code XYZ in genova"
		add "Fix response message in redecompras"

	Options:
		%s  Show help
	`,
		a.currentDaily.Date(),
		strings.Join(HelpFlags, ", "),
	)
}

func (a *add) Exec(args ...string) error {
	item := strings.Join(args, " ")
	item = strings.TrimSpace(item)

	if len(item) == 0 {
		return errors.ErrEmptyItem
	}

	if err := a.currentDaily.Add(item); err != nil {
		return err
	}

	fmt.Println(a.currentDaily.String())

	return nil
}
