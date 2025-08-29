package cmd

import "github.com/leoviggiano/daily/pkg/daily"

var HelpFlags = []string{"-h", "--help"}

type Command interface {
	Name() []string
	Help() string
	Exec(args ...string) error
}

func Commands(currentDaily *daily.Daily) []Command {
	return []Command{
		Add(currentDaily),
		Show(currentDaily),
		Delete(currentDaily),
		History(),
		Init(),
	}
}
