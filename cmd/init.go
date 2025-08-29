package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/leoviggiano/daily/pkg/daily"
)

func Init() Command {
	return &initDaily{}
}

type initDaily struct {
}

func (a *initDaily) Name() []string {
	return []string{"init"}
}

func (a *initDaily) Help() string {
	return fmt.Sprintf(`
	Começa a daily no dia de hoje: %s

	Usage:
		init

	Options:
		%s  Show help
	`,
		time.Now().Format("02-01-2006"),
		strings.Join(HelpFlags, ", "),
	)
}

func (a *initDaily) Exec(args ...string) error {
	daily := daily.NewDaily()

	if err := daily.Save(); err != nil {
		return err
	}

	fmt.Println("Daily iniciada com sucesso")
	fmt.Println(daily.String())

	return nil
}
