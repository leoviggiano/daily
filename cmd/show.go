package cmd

import (
	"fmt"
	"strings"

	"github.com/leoviggiano/daily/pkg/daily"
)

func Show(currentDaily *daily.Daily) Command {
	return &show{currentDaily: currentDaily}
}

type show struct {
	currentDaily *daily.Daily
}

func (s *show) Name() []string {
	return []string{"show", "s"}
}

func (s *show) Help() string {
	return fmt.Sprintf(`
	Mostra a daily atual: %s

	Usage:
		show

	Options:
		%s Show help
	`,
		s.currentDaily.Date(),
		strings.Join(HelpFlags, ", "),
	)
}

func (s *show) Exec(args ...string) error {
	fmt.Printf(`
--------------------------------
%s
--------------------------------
`,
		s.currentDaily.String(),
	)

	return nil
}
