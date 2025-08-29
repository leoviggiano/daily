package daily

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func historyPath() string {
	_, rootPath, _, _ := runtime.Caller(0)
	rootPath = strings.Split(rootPath, "/history.go")[0]
	rootPath = filepath.Join(rootPath, "history")

	return rootPath
}

func monthFolder(time time.Time) string {
	month := time.Month()
	year := time.Year()
	return strings.ToLower(fmt.Sprintf("%d-%02d", year, month))
}

func GetDaily(date time.Time) *Daily {
	history := GetHistory(date)

	for _, d := range history {
		if d.Created.Format("02-01-2006") == date.Format("02-01-2006") {
			return d
		}
	}

	return nil
}

func GetHistory(date time.Time) []*Daily {
	history := make([]*Daily, 0)

	rootPath := historyPath()
	rootPath = filepath.Join(rootPath, monthFolder(date))

	files, err := os.ReadDir(rootPath)
	if err != nil {
		return nil
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		path := path.Join(rootPath, file.Name())
		content, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err)
		}

		var d Daily
		err = json.Unmarshal(content, &d)
		if err != nil {
			log.Fatal(err)
		}

		history = append(history, &d)
	}

	return history
}

func GetAllDirectories() []string {
	months := make([]string, 0)

	rootPath := historyPath()

	files, err := os.ReadDir(rootPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.IsDir() {
			months = append(months, file.Name())
		}

	}

	return months
}

func getCurrentDaily() *Daily {
	now := time.Now()
	history := GetHistory(now)

	for _, d := range history {
		if d.Created.Format("02-01-2006") == now.Format("02-01-2006") {
			return d
		}
	}

	return nil
}
