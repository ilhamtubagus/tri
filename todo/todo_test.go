package todo

import (
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"sort"
	"testing"
)

func TestSetPriority(t *testing.T) {
	testPriorityTable := []struct {
		input    int
		expected int
	}{
		{
			input:    1,
			expected: 1,
		},
		{
			input:    3,
			expected: 3,
		},
		{
			input:    2,
			expected: 2,
		},
		{
			input:    4,
			expected: 2,
		},
	}

	for _, test := range testPriorityTable {
		item := Item{}
		item.SetPriority(test.input)

		assert.Equal(t, test.expected, item.Priority, "Expected priority %d, got %d", test.input, item.Priority)
	}
}

func TestPrettyP(t *testing.T) {
	testPriorityTable := []struct {
		input    int
		expected string
	}{
		{
			input:    1,
			expected: "(1)",
		},
		{
			input:    3,
			expected: "(3)",
		},
		{
			input:    2,
			expected: "",
		},
	}

	for _, test := range testPriorityTable {
		item := Item{}
		item.SetPriority(test.input)

		assert.Equal(t, test.expected, item.PrettyP(), "Expected priority %d, got %s", test.input, item.PrettyP())
	}
}

func TestLabel(t *testing.T) {
	item := Item{position: 1}

	assert.Equal(t, "1.", item.Label())
}

func TestPrettyStatus(t *testing.T) {
	testStatusTable := []struct {
		input    bool
		expected string
	}{
		{
			input:    true,
			expected: "[V]",
		},
		{
			input:    false,
			expected: "[ ]",
		},
	}

	for _, test := range testStatusTable {
		item := Item{}
		item.Done = test.input

		assert.Equal(t, test.expected, item.PrettyStatus(), "Expected status %t, got %s", test.input, item.PrettyStatus())
	}
}

func TestByPri(t *testing.T) {
	itemA := Item{Priority: 1, position: 3}
	itemB := Item{Priority: 2, position: 2}
	itemC := Item{Priority: 3, position: 1}
	itemD := Item{Priority: 3, position: 4}
	itemE := Item{Priority: 3, position: 5, Done: true}
	items := []Item{
		itemA,
		itemB,
		itemC,
		itemD,
		itemE,
	}

	sort.Sort(ByPri(items))

	assert.Equal(t, []Item{
		itemE,
		itemD,
		itemC,
		itemB,
		itemA,
	}, items)
}

func TestByPri_Len(t *testing.T) {
	items := ByPri{}

	assert.Equal(t, 0, items.Len())
}

func setupSuite(t *testing.T) func(t *testing.T) {
	mockOsMkdirAll := func(path string, perm os.FileMode) error {
		return nil
	}
	mockOsWriteFile := func(filename string, data []byte, perm os.FileMode) error {
		return nil
	}
	mockOsStat := func(name string) (os.FileInfo, error) {
		return nil, nil
	}
	mockOsReadFile := func(filename string) ([]byte, error) {
		return nil, nil
	}
	mockFilePathDir := func(path string) string {
		return ""
	}
	osMkdirAll = mockOsMkdirAll
	osWriteFile = mockOsWriteFile
	osStat = mockOsStat
	osReadFile = mockOsReadFile
	filepathDir = mockFilePathDir

	return func(t *testing.T) {
		osMkdirAll = mockOsMkdirAll
		osWriteFile = mockOsWriteFile
		osStat = mockOsStat
		osReadFile = mockOsReadFile
		filepathDir = mockFilePathDir
		jsonMarshal = json.Marshal
		jsonUnmarshal = json.Unmarshal
	}
}

func TestSaveItems_ErrJsonMarshal(t *testing.T) {
	teardown := setupSuite(t)
	defer teardown(t)
	mockError := errors.New("json marshal error")
	jsonMarshal = func(v interface{}) ([]byte, error) {
		return nil, mockError
	}

	err := SaveItems("filename", []Item{})

	assert.Equal(t, err, mockError)
}

func TestSaveItems_ErrOsMkdirAll(t *testing.T) {
	teardown := setupSuite(t)
	defer teardown(t)
	osMkdirAll = func(path string, perm os.FileMode) error {
		return os.ErrPermission
	}

	err := SaveItems("filename", []Item{})

	assert.Equal(t, err, os.ErrPermission)
}

func TestSaveItems_ErrOsWriteFile(t *testing.T) {
	teardown := setupSuite(t)
	defer teardown(t)
	osWriteFile = func(filename string, data []byte, perm os.FileMode) error {
		return os.ErrDeadlineExceeded
	}

	err := SaveItems("filename", []Item{})

	assert.Equal(t, err, os.ErrDeadlineExceeded)
}

func TestSaveItems(t *testing.T) {
	teardown := setupSuite(t)
	defer teardown(t)
	items := []Item{{Text: "test", Priority: 1}}

	err := SaveItems("filename", items)

	assert.Nil(t, err)
}
