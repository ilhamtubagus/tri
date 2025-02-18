package todo

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
)

var (
	osMkdirAll    = os.MkdirAll
	osWriteFile   = os.WriteFile
	osStat        = os.Stat
	osReadFile    = os.ReadFile
	filepathDir   = filepath.Dir
	jsonMarshal   = json.Marshal
	jsonUnmarshal = json.Unmarshal
)

type Item struct {
	Text     string `json:"text"`
	Priority int    `json:"priority"`
	position int
	Done     bool `json:"done"`
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

func (i *Item) PrettyStatus() string {
	if i.Done {
		return "[V]"
	}
	return "[ ]"
}

type ByPri []Item

func (b ByPri) Len() int {
	return len(b)
}

func (b ByPri) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByPri) Less(i, j int) bool {
	if b[i].Done != b[j].Done {
		return b[i].Done
	}

	if b[i].Priority != b[j].Priority {
		return b[i].Priority > b[j].Priority
	}

	return b[i].position > b[j].position
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
	b, err := jsonMarshal(items)
	if err != nil {
		return err
	}

	dir := filepathDir(filename)
	if err := osMkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	err = osWriteFile(filename, b, 0644)
	if err != nil {
		return err
	}

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
	_, err := osStat(filename)
	if err != nil {
		return nil, nil
	}

	b, err := osReadFile(filename)
	if err != nil {
		return nil, err
	}

	var items []Item
	if err := jsonUnmarshal(b, &items); err != nil {
		return nil, err
	}
	for i, _ := range items {
		items[i].position = i + 1
	}
	return items, nil
}
