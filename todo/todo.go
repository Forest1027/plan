package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

var JsonMarshal = json.Marshal

type Item struct {
	Text     string
	Priority int
}

func SaveItems(filename string, items []Item) error {
	b, err := JsonMarshal(items)
	if err != nil {
		return err
	}
	fmt.Println(string(b))

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadItems(filename string) ([]Item, error) {
	b, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, err
	}
	return items, err
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}
