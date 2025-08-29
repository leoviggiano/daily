package daily

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Item struct {
	Order       int    `json:"order"`
	Description string `json:"description"`
}

func (i *Item) String() string {
	return fmt.Sprintf("%d. %s", i.Order, i.Description)
}

type Daily struct {
	Created time.Time `json:"date"`
	List    []Item    `json:"list"`
}

func NewDaily() *Daily {
	if daily := getCurrentDaily(); daily != nil {
		return daily
	}

	return &Daily{
		Created: time.Now(),
		List:    []Item{},
	}
}

func (d *Daily) Add(item string) error {
	d.List = append(d.List, Item{
		Order:       len(d.List) + 1,
		Description: item,
	})

	return d.Save()
}

func (d *Daily) Remove(i []int) error {
	if len(i) == 0 {
		return nil
	}

	deleteMap := make(map[int]bool)
	for _, idx := range i {
		deleteMap[idx] = true
	}

	var newItems []Item
	for i, val := range d.List {
		if !deleteMap[i] {
			newItems = append(newItems, val)
		}
	}
	d.List = newItems

	return d.Save()
}

func (d *Daily) Date() string {
	return d.Created.Format("02-01-2006")
}

func (d *Daily) String() string {
	items := []string{}
	for _, item := range d.List {
		items = append(items, item.String())
	}

	return fmt.Sprintf("Daily: %s\n\n\t%s\n", d.Date(), strings.Join(items, "\n\t"))
}

func (d *Daily) json() []byte {
	json, err := json.Marshal(d)
	if err != nil {
		log.Fatal(err)
		return []byte{}
	}

	return json
}

func (d *Daily) Save() error {
	d.reoder()

	rootPath := historyPath()

	fileName := fmt.Sprintf("%s.txt", d.Created.Format("02-01-2006"))
	folder := filepath.Join(rootPath, monthFolder(d.Created))

	filePath := filepath.Join(folder, fileName)

	if _, err := os.Stat(folder); os.IsNotExist(err) {
		os.MkdirAll(folder, 0755)
	}

	return os.WriteFile(filePath, d.json(), 0644)
}

func (d *Daily) reoder() {
	for i, item := range d.List {
		d.List[i] = Item{
			Order:       i + 1,
			Description: item.Description,
		}
	}
}
