package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/nexidian/gocliselect"

	"github.com/leoviggiano/daily/pkg/daily"
)

func History() Command {
	return &history{}
}

type history struct {
}

func (s *history) Name() []string {
	return []string{"history", "h"}
}

func (s *history) Help() string {
	return fmt.Sprintf(`
	Mostra o histórico de dailies

	Usage:
		history
		history 29-08-2025

	Options:
		<date> Data da daily a ser exibida, no formato DD-MM-YYYY
		%s History help
	`,
		strings.Join(HelpFlags, ", "),
	)
}

func (s *history) Exec(args ...string) error {
	var daily *daily.Daily
	var err error

	if len(args) == 0 {
		daily, err = s.selectDaily()
	} else {
		daily, err = s.getDaily(args[0])
	}

	if err != nil {
		return err
	}

	fmt.Printf(`
--------------------------------
%s
--------------------------------
`,
		daily.String(),
	)

	return nil
}

func (s *history) getDaily(date string) (*daily.Daily, error) {
	parsedDate, err := time.Parse("02-01-2006", date)
	if err != nil {
		return nil, err
	}

	return daily.GetDaily(parsedDate), nil
}

func (s *history) selectDaily() (*daily.Daily, error) {
	history, err := s.selectHistory()
	if err != nil {
		return nil, err
	}

	menu := gocliselect.NewMenu("Selecione a daily")
	for _, d := range history {
		menu.AddItem(d.Created.Format("02-01-2006"), d.Created.Format("02-01-2006"))
	}

	selectedDaily := menu.Display()
	parsedDaily, err := time.Parse("02-01-2006", selectedDaily)
	if err != nil {
		return nil, err
	}

	return daily.GetDaily(parsedDaily), nil

}

func (s *history) selectHistory() ([]*daily.Daily, error) {
	months := daily.GetAllDirectories()

	menu := gocliselect.NewMenu("Selecione o mês")

	for _, month := range months {
		menu.AddItem(month, month)
	}

	selectedDate := menu.Display()

	parsedDate, err := time.Parse("2006-01", selectedDate)
	if err != nil {
		return nil, err
	}

	return daily.GetHistory(parsedDate), nil
}
