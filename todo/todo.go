package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

type Item struct {
	Text     string `json:"text"`
	Priority int    `json:"priority"`
	position int
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

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 3 {
		return "(3)"
	}

	return ""
}

func (i *Item) Label() string {
	return strconv.Itoa(i.position) + "."
}

type ByPri []Item

func (b ByPri) Len() int {
	return len(b)
}

func (b ByPri) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByPri) Less(i, j int) bool {
	if b[i].Priority == b[j].Priority {
		return b[i].position < b[j].position
	}

	return b[i].Priority > b[j].Priority
}

// SaveItems saves a list of todo items to a JSON file.
//
// Parameters:
//   - filename: The path to the file where the items will be saved.
//   - items: A slice of Item structs representing the todo items to be saved.
//
// Returns:
//   - error: An error if the items cannot be marshaled or the file cannot be written, nil otherwise.
func SaveItems(filename string, items []Item) error {
	b, err := json.Marshal(items)
	if err != nil {
		return err
	}

	dir := filepath.Dir(filename)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	err = os.WriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

	fmt.Println(string(b))

	return nil
}

// ReadItems reads and parses a list of todo items from a JSON file.
//
// Parameters:
//   - filename: The path to the JSON file containing the todo items.
//
// Returns:
//   - []Item: A slice of Item structs representing the todo items.
//   - error: An error if the file cannot be read or parsed, nil otherwise.

func ReadItems(filename string) ([]Item, error) {
	_, err := os.Stat(filename)
	if err != nil {
		return nil, nil
	}

	b, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var items []Item
	if err := json.Unmarshal(b, &items); err != nil {
		return nil, err
	}
	for i, _ := range items {
		items[i].position = i + 1
	}
	return items, nil
}
