package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nexidian/gocliselect"

	"github.com/leoviggiano/daily/pkg/daily"
	"github.com/leoviggiano/daily/pkg/errors"
)

func Delete(currentDaily *daily.Daily) Command {
	return &delete{currentDaily: currentDaily}
}

type delete struct {
	currentDaily *daily.Daily
}

func (d *delete) Name() []string {
	return []string{"delete", "d"}
}

func (d *delete) Help() string {
	return fmt.Sprintf(`
	Remove um item da daily atual: %s

	Usage:
		delete <item>

	Examples:
		delete 1
		d 2
		delete 1,2,3
		delete

	Options:
	    <item>  Item number to delete, must be a number or a comma separated list of numbers
		%s  Show help
	`,
		d.currentDaily.Date(),
		strings.Join(HelpFlags, ", "),
	)
}

func (d *delete) Exec(args ...string) error {
	items, err := d.getItems(args)
	if err != nil {
		return err
	}

	if err := d.currentDaily.Remove(items); err != nil {
		return err
	}

	fmt.Println(d.currentDaily.String())

	return nil
}

func (d *delete) getItems(args []string) ([]int, error) {
	if len(args) == 0 {
		args = d.menuInput()
	}

	items := make([]int, 0)
	for _, arg := range args {
		arg = strings.ReplaceAll(arg, ",", "")
		if arg == " " || arg == "" {
			continue
		}

		item, err := strconv.Atoi(arg)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", errors.ErrInvalidItem, arg)
		}

		// -1 because the item is 1-indexed
		items = append(items, item-1)
	}

	return items, nil
}

func (d *delete) menuInput() []string {
	items := d.currentDaily.List
	if len(items) == 0 {
		return []string{}
	}

	menu := gocliselect.NewMenu("Selecione o item para deletar")

	for _, item := range items {
		menu.AddItem(fmt.Sprintf("%d - %s", item.Order, item.Description), strconv.Itoa(item.Order))
	}

	return []string{menu.Display()}
}
