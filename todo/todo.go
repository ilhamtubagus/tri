package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Item struct {
	Text string `json:"text"`
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

	return items, nil
}
